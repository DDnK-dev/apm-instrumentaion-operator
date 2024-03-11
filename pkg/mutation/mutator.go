package mutation

import (
	"context"
	"errors"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/java"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

func NewPodMutator(pod *corev1.Pod, client client.Client) (*PodMutator, error) {
	mutators := []instrument.Mutator{
		java.Injector{}.SetClient(client),
	}

	return &PodMutator{
		pod:      pod,
		client:   client,
		mutators: mutators,
	}, nil
}

// PodMutator is a container for mutation logic.
type PodMutator struct {
	pod      *corev1.Pod
	client   client.Client
	mutators []instrument.Mutator
}

func (p PodMutator) Mutate(ctx context.Context) (*corev1.Pod, error) {
	if isAlreadyInstrumented(p.pod) {
		return p.pod, nil
	}

	var err error
	// check mutation condition
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, mutator := range p.mutators {
		p.pod, err = mutator.MutatePod(ctx, p.pod)
		if err != nil {
			return nil, err
		}
	}

	if !isValidInjection(p.pod) {
		return nil, errors.New("duplicate instrumentation exists, check spec")
	}
	return p.pod, nil
}

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

func isValidInjection(pod *corev1.Pod) bool {
	if !utils.HasAnnotation(pod.Annotations, consts.InstAnnotationInstrumented) {
		return false
	}
	defer func() {
		delete(pod.Annotations, consts.InstAnnotationInstrumented)
	}()
	containers := strings.Split(consts.InstAnnotationInstrumented, ", ")
	contMap := make(map[string]struct{})
	for _, cont := range containers {
		if _, exists := contMap[cont]; !exists {
			contMap[cont] = struct{}{}
			continue
		}
		return false
	}
	return true
}
