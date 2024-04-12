package python

import (
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/mutation/types"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

// verify javaInjector implements mutation.Injector
var _ types.Injector = &Injector{}

// Mutator implements mutation.Injector
// this is the struct that implements the Mutator interface
type Injector struct {
	Client client.Client
	Plan   *types.Plan
}

func NewInjector() types.Injector {
	return &Injector{}
}

func (i *Injector) PlanMutation(pod *corev1.Pod, annoMap utils.AnnotationMap) *types.Plan {
	p := &types.Plan{}
	value, ok := annoMap.GetAnnotationValue(consts.InstAnnotationKeyPython)
	if !ok {
		return p
	}
	inst, err := instrument.GetInstrumentFromAnnotation(value, i.Client)
	if err != nil {
		return p
	}
	containers, err := instrument.GetContainerNameFromAnnotation(pod, &annoMap, consts.InjAnnotationKeyPythonCont)
	if err != nil {
		return p
	}
	p.Instrumentation = inst
	p.Containers = containers
	return p
}

func (i *Injector) SetClient(c client.Client) types.Injector {
	i.Client = c
	return i
}

func (i *Injector) SetPlan(plan *types.Plan) {
	i.Plan = plan
}

func (i *Injector) GetPlan() *types.Plan {
	return i.Plan
}

func (i *Injector) Mutate(_ context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	var (
		pSpec   *v1.Python
		err     error
		contIdx int
	)
	pSpec, err = initPythonSpec(&i.Plan.Instrumentation.Spec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize java spec")
	}
	for _, cont := range i.Plan.Containers {
		contIdx = instrument.GetContainerIndex(pod, cont)
		if pod, err = injectPythonAgent(pSpec, pod, contIdx); err != nil {
			return nil, errors.Wrap(err, "failed to inject javaagent")
		}
	}
	return pod, nil
}
