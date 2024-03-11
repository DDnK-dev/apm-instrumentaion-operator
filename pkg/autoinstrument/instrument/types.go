package instrument

import (
	"context"
	"errors"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
)

type Mutator interface {
	MutatePod(ctx context.Context, pod *corev1.Pod) (*corev1.Pod, error)
	Name() string
	SetClient(cli client.Client) Mutator
	DirtyCheck(pod *corev1.Pod) error
	GetInstrumentation(ctx context.Context, pod *corev1.Pod) (*InstrumentationWithContainers, error)
	MarkMutatedContainers(pod *corev1.Pod, containers []string) *corev1.Pod
}

type InstrumentationWithContainers struct {
	Instrumentation v1.Instrumentation
	Containers      []string
}

var (
	ErrAlreadyInstrumented     = errors.New("pod is already instrumented")
	ErrInvalidAnnotationCont   = errors.New("duplicated target container annotation setting exists")
	ErrInvalidAnnotationValue  = errors.New("invalid annotation value")
	ErrInstrumentationNotFound = errors.New("instrumentation not found")
)
