package instrument

import (
	"context"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// verify javaInjector implements mutation.Mutator
var _ Mutator = &Injector{}

// Injector implements mutation.Mutator
// this is the struct that implements the Mutator interface
type Injector struct {
	client client.Client
}

// Name returns name of injector
// IMPLEMENT THIS
func (i Injector) Name() string {
	panic("Name not implemented")
	return ""
}

// SetClient set client used to get instrumentation or other k8s resources
func (i Injector) SetClient(cli client.Client) Mutator {
	i.client = cli
	return &i
}

// DirtyCheck check if pod already instrumented by other webhook or something
// If pod is already instrumented, returns error. should be called in MutatePod before mutate pod
// IMPLEMENT THIS
func (i Injector) DirtyCheck(_ *corev1.Pod) error {
	panic("DirtyCheck not instrumented")
	return nil
}

// MutatePod mutates pod using pod's annotation.
// IMPLEMENT THIS
func (i Injector) MutatePod(_ context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	panic("mutatedPod method didn't instrumented")
	return pod, nil
}

func (i Injector) GetInstrumentation(_ context.Context, pod *corev1.Pod) (*InstrumentationWithContainers, error) {
	panic("GetInstrumentation not instrumented")
	return nil, nil
}

// MarkMutatedContainers mark containers that instrumented containers
func (i Injector) MarkMutatedContainers(pod *corev1.Pod, containers []string) *corev1.Pod {
	if len(containers) == 0 {
		return pod
	}
	// flattening string slice to string to cont0, cont1, cont2
	var annonValue string
	for _, cont := range containers {
		annonValue += cont + ","
	}
	pod.Annotations[consts.InstAnnotationInstrumented] = annonValue[:len(annonValue)-1] // remove last comma
	return pod
}
