/*
this package contains all the constants used in the project
*/

package consts

import "errors"

const (
	domain = "apm.ogas.kr"
)

// Annotation Keys
const (
	// instInjectKey is the annotation key used to specify the containers that have to be instrumented
	instInjectKey                = domain + "/inject"
	InstAnnotationKeyJava        = instInjectKey + "-java"
	InstAnnotationKeyGo          = instInjectKey + "-go"
	InstAnnotationKeyPython      = instInjectKey + "-python"
	InstAnnotationKeyDotNet      = instInjectKey + "-dotnet"
	InstAnnotationKeyNodeJS      = instInjectKey + "-nodejs"
	InstAnnotationKeyApacheHttpd = instInjectKey + "-apache-httpd"
	InstAnnotationKeyNginx       = instInjectKey + "-nginx"
	InstAnnotationKeySDK         = instInjectKey + "-sdk"

	// InstAnnotationKeyContName is annotation key used to specify which containers should be instrumented.
	InstAnnotationKeyContName  = domain + "/container-names"
	InstAnnotationKeyJavaCont  = domain + "/java-container-names"
	InjAnnotationKeyGoCont     = domain + "/go-container-names"
	InjAnnotationKeyPythonCont = domain + "/python-container-names"
	InjAnnotationKeyDotNetCont = domain + "/dotnet-container-names"
	InjAnnotationKeyNodeJSCont = domain + "/nodejs-container-names"
	InjAnnotationKeyApacheCont = domain + "/apache-httpd-container-names"
	InjAnnotationKeyNginxCont  = domain + "/nginx-container-names"
	InjAnnotationKeySDKCont    = domain + "/sdk-container-names"

	InstAnnotationKeyGoTarget = domain + "/go-auto-target-exe"

	InstAnnotationKeyDotNetRuntime = domain + "/dotnet-auto-runtime"
)

// InitContainer Constants
const (
	APMInitContainerName       = "ogas-instrument-otel"
	APMInitContainerNameJava   = APMInitContainerName + "-java"
	APMInitContainerNameGo     = APMInitContainerName + "-go"
	APMInitContainerNamePython = APMInitContainerName + "-python"
	APMInitContainerNameDotNet = APMInitContainerName + "-dotnet"

	APMVolumeName       = "ogas-instrument"
	APMVolumeNameJava   = APMVolumeName + "-java"
	APMVolumeNamePython = APMVolumeName + "-python"
)

// Image Constants
const (
	APMImageJava   = "otel/autoinstrumentation-java:latest"
	APMImagePython = "otel/autoinstrumentation-python:latest"
)

// Java Configs
const (
	LoggingSimple = "simple"
)

// Common Trace Environment Variables
const (
	EnvExporterEndpoint                = "OTEL_EXPORTER_OTLP_ENDPOINT"
	EnvTraceSampler                    = "OTEL_TRACES_SAMPLER"
	EnvTraceSamplerArg                 = "OTEL_TRACES_SAMPLER_ARG"
	EnvPropagators                     = "OTEL_PROPAGATORS"
	EnvServiceName                     = "OTEL_SERVICE_NAME"
	EnvTracesExporter                  = "OTEL_TRACES_EXPORTER"
	EnvOtelExporterOTLPTracesProtocol  = "OTEL_EXPORTER_OTLP_TRACES_PROTOCOL"
	EnvMetricsExporter                 = "OTEL_METRICS_EXPORTER"
	EnvOtelExporterOTLPMetricsProtocol = "OTEL_EXPORTER_OTLP_METRICS_PROTOCOL"
	EnvLogsExporter                    = "OTEL_LOGS_EXPORTER"
	EnvJavaAgentLogging                = "OTEL_JAVAAGENT_LOGGING"
)

// Predefined Errors
var (
	ErrNotDefined = errors.New("not defined")
	ErrNotValid   = errors.New("not valid value")
)
