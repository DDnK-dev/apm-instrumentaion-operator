package python

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"

	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/autoinstrument/instrument"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
)

const (
	pythonPathEnv            = "PYTHONPATH"
	pythonPathPrefix         = "/otel-auto-instrumentation-python/opentelemetry/instrumentation/auto_instrumentation"
	pythonPathPostfix        = "/otel-auto-instrumentation-python"
	pythonContainerMountPath = pythonPathPostfix
	volumeName               = consts.APMVolumeNamePython
	pythonInitContainerName  = consts.APMInitContainerNamePython

	defaultTraceExporter      = "otlp"
	defaultOtelTraceProtocol  = "http/protobuf"
	defaultMetricsExporter    = "otlp"
	defaultOtelMetricProtocol = "http/protobuf"
)

func initPythonSpec(spec *v1.InstrumentationSpec) (*v1.Python, error) {
	var (
		pSpec = spec.Python
	)
	if pSpec.Endpoint == "" {
		pSpec.Endpoint = spec.Endpoint
	}
	// sampling setting and validation
	if pSpec.Sampling.Sampler == "" && pSpec.Sampling.SamplerArg == "" {
		pSpec.Sampling = spec.Sampling
	}
	conf, err := instrument.OverrideConfiguration(&spec.Configuration, &pSpec.Config)
	if err != nil {
		return nil, err
	}
	pSpec.Config = *conf
	return &pSpec, nil
}

func injectPythonAgent(pSpec *v1.Python, pod *corev1.Pod, index int) (*corev1.Pod, error) {
	var (
		idx       int
		container *corev1.Container
		err       error
	)
	// validate if pythonPathEnv exists and use value from syntax
	container = &pod.Spec.Containers[index]
	idx, err = instrument.ValidateContainerEnv(container.Env, pythonPathEnv)
	if err != nil {
		return pod, err
	}
	// autoinstrumentation package should be prepended to PYTHONPATH to use the site module importing other integrations
	if idx == -1 { // if there is no pythonPathEnv env key
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  pythonPathEnv,
			Value: fmt.Sprintf("%s:%s", pythonPathPrefix, pythonPathPostfix),
		})
	} else { // if already exists, append agent env var
		container.Env[idx].Value =
			fmt.Sprintf("%s:%s:%s", pythonPathPrefix, container.Env[idx].Value, pythonPathPostfix)
	}

	// if endpoint is set to 4317, which is grpc port, replace endpoint to 4318 and set protocol to HTTP
	// because default python agent uses HTTP protocol
	if idx = instrument.GetEnvVarIndex(container, consts.EnvExporterEndpoint); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvExporterEndpoint,
			Value: negotiateProtocol(pSpec),
		})
	}

	// set sampling
	if idx = instrument.GetEnvVarIndex(container, consts.EnvTraceSampler); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvTraceSampler,
			Value: pSpec.Sampling.Sampler,
		})
	}
	if idx = instrument.GetEnvVarIndex(container, consts.EnvTraceSamplerArg); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvTraceSamplerArg,
			Value: pSpec.Sampling.SamplerArg,
		})
	}

	// set tracer
	if idx = instrument.GetEnvVarIndex(container, consts.EnvTracesExporter); idx == -1 {
		tracer := pSpec.Config.Tracer
		if tracer != "none" {
			tracer = defaultTraceExporter
		}
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvTracesExporter,
			Value: tracer,
		})
	}
	if idx = instrument.GetEnvVarIndex(container, consts.EnvOtelExporterOTLPTracesProtocol); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvOtelExporterOTLPTracesProtocol,
			Value: defaultOtelTraceProtocol,
		})
	}

	// set serviceNameLabel
	if idx = instrument.GetEnvVarIndex(container, consts.EnvServiceName); idx == -1 {
		var serviceName string
		if serviceName = pod.Labels[pSpec.Config.ServiceNameLabel]; serviceName == "" {
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
			Value: strings.Join(pSpec.Config.Propagator, ","),
		})
	}

	// set metrics
	if idx = instrument.GetEnvVarIndex(container, consts.EnvMetricsExporter); idx == -1 {
		metric := pSpec.Config.Metrics
		if metric != "none" {
			metric = defaultMetricsExporter
		}
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvMetricsExporter,
			Value: metric,
		})
	}
	if idx = instrument.GetEnvVarIndex(container, consts.EnvOtelExporterOTLPMetricsProtocol); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvOtelExporterOTLPMetricsProtocol,
			Value: defaultOtelMetricProtocol,
		})
	}

	// set logs
	if idx = instrument.GetEnvVarIndex(container, consts.EnvLogsExporter); idx == -1 {
		container.Env = append(container.Env, corev1.EnvVar{
			Name:  consts.EnvLogsExporter,
			Value: pSpec.Config.Logs,
		})
	}

	// inject Python instrumentation spec env vars.
	// if there is already an env var with the same name, it will be skipped.
	if len(pSpec.Config.EnvVars) > 0 {
		for _, envVar := range pSpec.Config.EnvVars {
			if idx = instrument.GetEnvVarIndex(container, envVar.Name); idx == -1 {
				container.Env = append(container.Env, envVar)
			}
		}
	}

	container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
		Name:      volumeName,
		MountPath: pythonContainerMountPath,
	})
	if !instrument.HasInitContainer(pod, pythonInitContainerName) {
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{
			Name: volumeName,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
		pod.Spec.InitContainers = append(pod.Spec.InitContainers, corev1.Container{
			Name:    pythonInitContainerName,
			Image:   pSpec.Image,
			Command: []string{"cp", "-r", "/autoinstrumentation/.", pythonContainerMountPath},
			VolumeMounts: []corev1.VolumeMount{{
				Name:      volumeName,
				MountPath: pythonContainerMountPath,
			}},
		})
	}
	return pod, nil
}

// negotiateProtocol changes endpoint port from `4317` to `4318` if needed
// because python autoinstrumentation used http/protobuf for default option.
func negotiateProtocol(pSpec *v1.Python) string {
	var (
		ep       string
		epTokens []string
		epLen    int
	)
	ep = pSpec.Endpoint
	epTokens = strings.Split(ep, ":")
	epLen = len(epTokens)

	if epLen != 0 && epTokens[epLen-1] == "4317" {
		epTokens[epLen-1] = "4318"
		ep = strings.Join(epTokens, ":")
	}
	return ep
}
