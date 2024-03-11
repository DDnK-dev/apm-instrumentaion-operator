package instrument

import (
	"errors"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

// ParseInstAnnotation parses annotation value and returns instrumentation name and namespace
// get pod namespace and annotation value as parameter
// returns target instrumentation name and namespace, error if exists
func ParseInstAnnotation(podNs string, annotation string) (instNS string, instName string, err error) {
	// case 1. if true,
	if annotation == "true" {
		return podNs, "", nil
	}

	src := strings.Split(annotation, "/")
	if len(src) == 1 {
		return podNs, src[0], nil
	} else if len(src) == 2 {
		return src[0], src[1], nil
	}
	return "", "", ErrInvalidAnnotationValue
}

func HasValidAnnotation(annotation map[string]string, key string) bool {
	if !utils.HasAnnotation(annotation, key) || annotation[key] == "false" {
		return false
	}
	return true
}

func GetEnvVarIndex(container corev1.Container, envKey string) (index int) {
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
