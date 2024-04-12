package java

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
)

const (
	envJavaToolsOptions   = "JAVA_TOOL_OPTIONS"
	javaJVMArgument       = " -javaagent:/otel-auto-instrumentation-java/javaagent.jar"
	javaInstrMountPath    = "/otel-auto-instrumentation-java"
	javaInitContainerName = consts.APMInitContainerNameJava
	javaVolumeName        = consts.APMVolumeNameJava
	javaVolumeLimit       = "200Mi"
)

// initJavaSpec initializes Java instrumentation spec from the given spec
func initJavaSpec(spec *v1.InstrumentationSpec) (*v1.Java, error) {
	var (
		jSpec = spec.Java
	)
	if jSpec.Endpoint == "" {
		jSpec.Endpoint = spec.Endpoint
	}
	// sampling setting and validation
	if jSpec.Sampling.Sampler == "" && jSpec.Sampling.SamplerArg == "" {
		jSpec.Sampling = spec.Sampling
	}
	conf, err := instrument.OverrideConfiguration(&spec.Configuration, &jSpec.Config)
	if err != nil {
		return nil, err
	}
	jSpec.Config = *conf
	return &jSpec, nil
}

// TODO: check; if the injection process is halted by the error, pod shouldn't be modified.
func injectJavaagent(jSpec *v1.Java, pod *corev1.Pod, index int) (*corev1.Pod, error) {
	var (
		idx       int
		container *corev1.Container
		err       error
	)
	// validateContainerEnv
	// if JAVA_TOOL_OPTIONS has valuesFrom, we can't inject javaagent.
	container = &pod.Spec.Containers[index]
	idx, err = instrument.ValidateContainerEnv(container.Env, envJavaToolsOptions)
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

	// set endpoint
	// check if the endpoint is already set
	if idx = instrument.GetEnvVarIndex(container, consts.EnvExporterEndpoint); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvExporterEndpoint,
			Value: jSpec.Endpoint,
		})
	}
	// set sampling
	if idx = instrument.GetEnvVarIndex(container, consts.EnvTraceSampler); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvTraceSampler,
			Value: jSpec.Sampling.Sampler,
		})
	}
	if idx = instrument.GetEnvVarIndex(container, consts.EnvTraceSamplerArg); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvTraceSamplerArg,
			Value: jSpec.Sampling.SamplerArg,
		})
	}

	// set java logging
	if idx = instrument.GetEnvVarIndex(container, consts.EnvJavaAgentLogging); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvJavaAgentLogging,
			Value: jSpec.Logging,
		})
	}

	// set Configuration
	// set tracer
	if idx = instrument.GetEnvVarIndex(container, consts.EnvTracesExporter); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvTracesExporter,
			Value: jSpec.Config.Tracer,
		})
	}

	// set serviceNameLabel
	if idx = instrument.GetEnvVarIndex(container, consts.EnvServiceName); idx == -1 {
		var serviceName string
		if serviceName = pod.Labels[jSpec.Config.ServiceNameLabel]; serviceName == "" {
			if pod.Name == "" {
				serviceName = pod.GenerateName
			} else {
				serviceName = pod.Name
			}
		}
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvServiceName,
			Value: serviceName,
		})
	}

	// set propagator
	if idx = instrument.GetEnvVarIndex(container, consts.EnvPropagators); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvPropagators,
			Value: strings.Join(jSpec.Config.Propagator, ","),
		})
	}

	// set metrics
	if idx = instrument.GetEnvVarIndex(container, consts.EnvMetricsExporter); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvMetricsExporter,
			Value: jSpec.Config.Metrics,
		})
	}
	// set logs
	if idx = instrument.GetEnvVarIndex(container, consts.EnvLogsExporter); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvLogsExporter,
			Value: jSpec.Config.Logs,
		})
	}
	// inject Java instrumentation spec env vars.
	// if there is already an env var with the same name, it will be skipped
	for _, env := range jSpec.Config.EnvVars {
		if idx = instrument.GetEnvVarIndex(&pod.Spec.Containers[index], env.Name); idx == -1 {
			pod.Spec.Containers[index].Env = append(pod.Spec.Containers[index].Env, env)
		}
	}

	// create new volume mount
	container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
		Name:      javaVolumeName,
		MountPath: javaInstrMountPath,
	})
	// inject volumes and init containers for the first processed container
	if !instrument.HasInitContainer(pod, javaInitContainerName) {
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
