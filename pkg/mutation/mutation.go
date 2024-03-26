/*
Package mutation contains the logic for mutating the pod.
*/

package mutation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// validate if PodHandler implements admission.Handler
var _ admission.Handler = &PodHandler{}

// nolint:lll
//+kubebuilder:webhook:path=/mutate-v1-pod,mutating=true,failurePolicy=ignore,groups="",resources=pods,verbs=create,versions=v1,name=mpod.kb.io,admissionReviewVersions=v1,sideEffects=None
//+kubebuilder:rbac:groups="",resources=namespaces,verbs=list;watch

// PodHandler implements admission.DecoderInjector.
type PodHandler struct {
	Client   client.Client
	Decoder  *admission.Decoder
	Recorder record.EventRecorder
}

// Handle function handles admission request from kubernetes api server
// and check if pod is mutation target, then validates, process pod mutation
func (p *PodHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	var (
		pod     = &corev1.Pod{}
		mutator = &PodMutator{}
	)

	err := p.Decoder.Decode(req, pod)
	if err != nil {
		p.Recorder.Event(pod, "Warning", "Decoding",
			fmt.Sprintf("Failed to decode pod from request becaues of error %v", err))
		return admission.Errored(http.StatusBadRequest, err)
	}

	// check if pod's namespace is not created yet
	ns := corev1.Namespace{}
	err = p.Client.Get(ctx, types.NamespacedName{Name: req.Namespace, Namespace: ""}, &ns)
	if err != nil {
		res := admission.Errored(http.StatusInternalServerError, err)
		res.Allowed = true // set error, but doesn't block pod creation
		p.Recorder.Event(pod, "Normal", "Skipped",
			"pod instrumentation passed because pod created before namespace creation")
		return res
	}

	if mutator, err = NewPodMutator(pod, p.Client); err != nil {
		p.Recorder.Event(pod, "Warning", "Mutator",
			fmt.Sprintf("Failed to create pod mutator becaues of error %v", err))
		return admission.Errored(http.StatusInternalServerError, err)
	}
	pod, err = mutator.Mutate(ctx)
	if err != nil {
		p.Recorder.Event(pod, "Normal", "Mutating",
			fmt.Sprintf("Failed to mutate pod becaus of error %v", err))
		return admission.Errored(http.StatusInternalServerError, err)
	}

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		p.Recorder.Event(pod, "Warning", "Marshal",
			fmt.Sprintf("Failed to malshal mutated pod becaus of error %v", err))
		return admission.Errored(http.StatusInternalServerError, err)
	}
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}
