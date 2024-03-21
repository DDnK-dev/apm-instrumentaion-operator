package java

import (
	"context"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/mutation/type"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

// verify javaInjector implements mutation.Injector
var _ _type.Injector = &Injector{}

// Mutator implements mutation.Injector
// this is the struct that implements the Mutator interface
type Injector struct {
	Client client.Client
	Plan   *_type.Plan
}

func NewInjector() _type.Injector {
	return &Injector{}
}

// TODO: need error handling and logging
func (i *Injector) PlanMutation(pod *corev1.Pod, labelMap utils.LabelMap) *_type.Plan {
	p := &_type.Plan{}
	value, ok := labelMap.GetLabelValue(consts.InstAnnotationKeyJava)
	if !ok {
		return p
	}
	inst, err := instrument.GetInstrumentFromAnnotation(value, i.Client)
	if err != nil {
		return nil
	}
	containers, err := instrument.GetContainerNameFromAnnotation(pod, &labelMap, consts.InstAnnotationKeyJavaCont)
	if err != nil {
		return nil
	}
	p.Instrumentation = inst
	p.Containers = containers
	return p
}

func (i *Injector) SetClient(c client.Client) _type.Injector {
	i.Client = c
	return i
}

func (i *Injector) SetPlan(plan *_type.Plan) {
	i.Plan = plan
}

func (i *Injector) GetPlan() *_type.Plan {
	return i.Plan
}

func (i *Injector) Mutate(_ context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	var (
		jSpec   *v1.Java
		err     error
		contIdx int
	)
	jSpec, err = initJavaSpec(&i.Plan.Instrumentation.Spec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize java spec")
	}
	for _, cont := range i.Plan.Containers {
		contIdx = instrument.GetContainerIndex(pod, cont)
		if pod, err = injectJavaagent(jSpec, pod, contIdx); err != nil {
			return nil, errors.Wrap(err, "failed to inject javaagent")
		}
	}
	return pod, nil
}
