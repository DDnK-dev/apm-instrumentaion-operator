package instrument

import (
	"context"
	"errors"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

// GetContainerIndex returns container index from pod. Index indicates the position of the container in the pod.
// return if the container does not exist
func GetContainerIndex(pod *corev1.Pod, container string) int {
	for i, c := range pod.Spec.Containers {
		if c.Name == container {
			return i
		}
	}
	return -1
}

func GetEnvVarIndex(container *corev1.Container, envKey string) (index int) {
	for i, env := range container.Env {
		if env.Name == envKey {
			return i
		}
	}
	return -1
}

// ValidateContainerEnv validates if envKey exists in container env and if it has valuesFrom field
func ValidateContainerEnv(envs []corev1.EnvVar, envKey string) (index int, err error) {
	for i := range envs {
		if envs[i].Name == envKey {
			if envs[i].ValueFrom != nil {
				return i, errors.New("env var has valuesFrom field")
			}
			return i, nil
		}
	}
	return -1, nil
}

// GetInstrumentFromAnnotation validates annotation value and returns instrumentation
func GetInstrumentFromAnnotation(value string, cli client.Client) (*v1.Instrumentation, error) {
	// validate annotation value
	// 1. if false, return nil
	if value == "false" {
		return nil, nil
	}
	// 2. if true
	// 	 list all instrumentations from namespace
	//   if  len(instrumentation) > 1 or 0, return nil
	if value == "true" {
		// get all instrumentations from namespace
		instList := &v1.InstrumentationList{}
		if err := cli.List(context.TODO(), instList); err != nil { // TODO: context.TODO() -> context something
			return nil, errors.New("failed to get instrumentation list")
		} else if len(instList.Items) != 1 {
			return nil, errors.New("instrumentation needs to be one on true option")
		}
		return &instList.Items[0], nil
	}
	// 3. if string slice splited by / is not 2, return nil
	//    if my-ns/my-inst format, get my-inst from my-ns
	values := strings.Split(value, "/")
	if len(values) != 2 {
		return nil, errors.New("invalid annotation value format (ns/inst)")
	}
	inst := &v1.Instrumentation{}
	if err := cli.Get(context.TODO(), types.NamespacedName{Namespace: values[0], Name: values[1]}, inst); err != nil {
		return nil, errors.New("instrumentation not found")
	}
	return inst, nil
}

// GetContainerNameFromAnnotation returns container names from annotation value
func GetContainerNameFromAnnotation(pod *corev1.Pod, l *utils.AnnotationMap, key string) ([]string, error) {
	// 1. if container-names exists, check container names are valid and return values (ignores key)
	containerNames, ok := l.GetAnnotationValue(consts.InstAnnotationKeyContName)
	if ok && containerNames != "" {
		return CheckContainerName(pod, containerNames)
	}
	// 2. if container-names does not exist, and annotation with key exists,
	// check if container names are valid and return values
	containerNames, ok = l.GetAnnotationValue(key)
	if ok && containerNames != "" {
		return CheckContainerName(pod, containerNames)
	}
	// if neither exists, return first container name
	return []string{pod.Spec.Containers[0].Name}, nil
}

func CheckContainerName(pod *corev1.Pod, containerNames string) ([]string, error) {
	containers := strings.Split(containerNames, ", ")
	for _, container := range containers {
		if HasContainer(pod, container) {
			return nil, errors.New("target container not found")
		}
	}
	return containers, nil
}

func HasContainer(pod *corev1.Pod, container string) bool {
	for _, c := range pod.Spec.Containers {
		if c.Name == container {
			return true
		}
	}
	return false
}

func OverrideConfiguration(base *v1.Configuration, lang *v1.Configuration) (*v1.Configuration, error) {
	if base == nil || lang == nil {
		return nil, errors.New("configuration is nil")
	}
	if lang.Tracer == "" {
		lang.Tracer = base.Tracer
	}
	if lang.ServiceNameLabel == "" {
		lang.ServiceNameLabel = base.ServiceNameLabel
	}
	if len(lang.Propagator) == 0 {
		lang.Propagator = base.Propagator
	}
	if len(lang.EnvVars) == 0 {
		lang.EnvVars = base.EnvVars
	}
	if lang.Metrics == "" {
		lang.Metrics = base.Metrics
	}
	if lang.Logs == "" {
		lang.Logs = base.Logs
	}
	return lang, nil
}

func HasInitContainer(pod *corev1.Pod, name string) bool {
	for _, c := range pod.Spec.InitContainers {
		if c.Name == name {
			return true
		}
	}
	return false
}
