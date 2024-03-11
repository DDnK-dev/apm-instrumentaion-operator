package java

import (
	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
)

const (
	envJavaToolsOptions   = "JAVA_TOOL_OPTIONS"
	javaJVMArgument       = " -javaagent:/otel-auto-instrumentation-java/javaagent.jar"
	javaInstrMountPath    = "/otel-auto-instrumentation-java"
	javaInitContainerName = consts.APMInitContainerNameJava
	javaVolumeName        = consts.APMVolumeName + "-java"
	javaVolumeLimit       = "200Mi"
)

func injectCommonSettings(spec *v1.InstrumentationSpec, pod *corev1.Pod, index int) (*corev1.Pod, error) {
	container = &pod.Spec.Containers[index]

	container = injectConfig(spec.Config, container)

	// set endpoint to pod
	// 일단 걍 넣는거임ㅎㅎ;;

	//container.Env = append(container.Env, corev1.EnvVar{
	//	Name:  "OTEL_EXPORTER_OTLP_ENDPOINT",
	//	Value: spec.Endpoint,
	//}
	// endpoint
	// sampler
	// sampler arg
	// tracer
	// serviceNameLabel
	// propagator
	// metrics
	// logs

	return pod, nil
}

func injectConfig(config *v1.Configuration, container *corev1.Container) *corev1.Container {
	if config.Endpoint != "" {

	}

	return container
}

func injectJavaagent(jSpec *v1.Java, pod *corev1.Pod, index int) (*corev1.Pod, error) {
	// validateContainerEnv
	// if JAVA_TOOL_OPTIONS has valuesFrom, we can't inject javaagent.
	container := &pod.Spec.Containers[index]
	{
		idx, err := instrument.ValidateContainerEnv(container.Env, envJavaToolsOptions)
		if err != nil {
			return pod, err
		}
		if idx == -1 { // if there is no javaTools Options env key
			container.Env = append(container.Env, corev1.EnvVar{
				Name:  envJavaToolsOptions,
				Value: javaJVMArgument,
			})
		} else {
			container.Env[idx].Value = container.Env[idx].Value + javaJVMArgument
		}
	}
	// inject Java instrumentation spec env vars.
	// if there is already an env var with the same name, it will be skipped
	for _, env := range jSpec.EnvVars {
		if idx := instrument.GetEnvVarIndex(pod.Spec.Containers[index], env.Name); idx == -1 {
			pod.Spec.Containers[index].Env = append(pod.Spec.Containers[index].Env, env)
		}
	}
	// create new volume mount
	container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
		Name:      javaVolumeName,
		MountPath: javaInstrMountPath,
	})
	// inject volumes and init containers for the first processed container
	if !utils.HasInitContainer(pod, javaInitContainerName) {
		volumeSizeLimit := resource.MustParse(javaVolumeLimit)
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{
			Name: javaVolumeName,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{
					SizeLimit: &volumeSizeLimit,
				},
			}})
		pod.Spec.InitContainers = append(pod.Spec.InitContainers, corev1.Container{
			Name:    javaInitContainerName,
			Image:   jSpec.Image,
			Command: []string{"cp", "/javaagent.jar", javaInstrMountPath + "/javaagent.jar"},
			VolumeMounts: []corev1.VolumeMount{{
				Name:      javaVolumeName,
				MountPath: javaInstrMountPath,
			}},
		})
	}
	return pod, nil
}
