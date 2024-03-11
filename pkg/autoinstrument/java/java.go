package java

import (
	"context"
	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
	"strings"

	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

// verify javaInjector implements mutation.Mutator
var _ instrument.Mutator = &Injector{}

// Injector implements mutation.Mutator
// this is the struct that implements the Mutator interface
type Injector struct {
	instrument.Injector
	Client client.Client
}

func (j Injector) Name() string {
	return "java"
}

func (j Injector) DirtyCheck(pod *corev1.Pod) error {
	if utils.HasAnnotation(pod.Annotations, consts.InstAnnotationKeyJava) {
		return instrument.ErrAlreadyInstrumented
	}
	return nil
}

func (j Injector) MutatePod(ctx context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	if j.DirtyCheck(pod) != nil {
		return pod, nil
	} else if !instrument.HasValidAnnotation(pod.Annotations, consts.InstAnnotationKeyJava) {
		return pod, nil
	}

	inst, err := j.GetInstrumentation(ctx, pod)
	if err != nil {
		return pod, err
	}

	for _, cont := range inst.Containers {
		index := utils.GetContIndexWithName(cont, pod)
		if index == -1 {
			return pod, errors.New("container not found")
		}
		// container validation check
		if len(pod.Spec.Containers) < index || index < 0 {
			return pod, errors.New("container index is not valid, container count is" +
				strconv.Itoa(len(pod.Spec.Containers)) + "index is" + strconv.Itoa(index))
		}
		// instrument container with common settings
		pod, err = injectCommonSettings(&inst.Instrumentation.Spec, pod, index)

		// instrument container with partial settings

		// instrument container with java settings
		pod, err = injectJavaagent(&inst.Instrumentation.Spec.Java, pod, index)
		if err != nil {
			return pod, err
		} else {
			// inject common settings
			// inject common sdk configuration
			// inject common env variables // -> inject common env var
			// inject security context (if needed)
		}
	}
	return pod, nil
}

// GetInstrumentation returns instrumentation InstrumentationWithContainers
// this includes the instrumentation and the containers to be instrumented
func (j Injector) GetInstrumentation(ctx context.Context, pod *corev1.Pod) (*instrument.InstrumentationWithContainers, error) {
	var res instrument.InstrumentationWithContainers

	// containerName checking
	hasContName := utils.HasAnnotation(pod.Annotations, consts.InstAnnotationKeyContName)
	hasJavaCont := utils.HasAnnotation(pod.Annotations, consts.InstAnnotationKeyJavaCont)
	if hasContName && hasJavaCont {
		return nil, instrument.ErrInvalidAnnotationCont
	} else if !hasContName && !hasJavaCont {
		res.Containers = append(res.Containers, pod.Spec.Containers[0].Name)
	} else if hasContName {
		res.Containers = strings.Split(pod.Annotations[consts.InstAnnotationKeyContName], ", ")
	} else if hasJavaCont {
		res.Containers = strings.Split(pod.Annotations[consts.InstAnnotationKeyJavaCont], ", ")
	}

	// get instrumentation
	if pod.Annotations[consts.InstAnnotationKeyJava] == "true" {
		instList := &v1.InstrumentationList{}
		if err := j.Client.List(ctx, instList); err != nil {
			return nil, errors.Wrap(err, "list instrumentation")
		}
		if len(instList.Items) > 1 {
			return nil, errors.New("annotation: true but multiple instrumentation exists")
		}
		res.Instrumentation = instList.Items[0]
	} else {
		ns, name, err := instrument.ParseInstAnnotation(pod.Namespace, pod.Annotations[consts.InstAnnotationKeyJava])
		if err != nil {
			return nil, errors.Wrapf(err, "annotation: %s", consts.InstAnnotationKeyJava)
		}
		if err = j.Client.Get(ctx, client.ObjectKey{Name: name, Namespace: ns}, &res.Instrumentation); err != nil {
			return nil, instrument.ErrInstrumentationNotFound
		}
	}
	return &res, nil
}
