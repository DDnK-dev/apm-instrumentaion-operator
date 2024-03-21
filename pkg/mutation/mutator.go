package mutation

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/java"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/mutation/type"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

// PodMutator is a container for mutation logic.
type PodMutator struct {
	pod      *corev1.Pod
	client   client.Client
	mutators []_type.Injector
}

func NewPodMutator(pod *corev1.Pod, client client.Client) (*PodMutator, error) {
	mutators := []_type.Injector{
		java.NewInjector(),
	}
	return &PodMutator{
		pod:      pod,
		client:   client,
		mutators: mutators,
	}, nil
}

func (p *PodMutator) Mutate(ctx context.Context) (*corev1.Pod, error) {
	if isAlreadyInstrumented(p.pod) {
		return p.pod, nil
	}

	// inject label validation -> if there is no inject label, we need to skip the pod
	lMap := utils.NewLabelMap(p.pod)
	for _, mutator := range p.mutators {
		mutator.SetClient(p.client).SetPlan(mutator.PlanMutation(p.pod, lMap))
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	for _, mutator := range p.mutators {
		if checkMutatorActive(mutator) {
			p.pod, err = mutator.Mutate(ctx, p.pod)
			if err != nil {
				return nil, err
			}
		}
	}
	return p.pod, nil
}

// isAlreadyInstrumented checks if the pod is already instrumented with init container name
func isAlreadyInstrumented(pod *corev1.Pod) bool {
	initContainerSet := map[string]struct{}{
		consts.APMInitContainerNameJava:   {},
		consts.APMInitContainerNameGo:     {},
		consts.APMInitContainerNamePython: {},
		consts.APMInitContainerNameDotNet: {},
	}
	for _, container := range pod.Spec.InitContainers {
		if _, ok := initContainerSet[container.Name]; ok {
			return true
		}
	}
	return false
}

func checkMutatorActive(mutator _type.Injector) bool {
	return mutator.GetPlan().Instrumentation != nil && mutator.GetPlan().Containers != nil
}
