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
	// InstAnnotationInstrumented is annotation key used to specify that which container is instrumented
	// At the end of mutation, this annotation will be used to check if there is duplicated mutation
	// As the limit of annotation value is 256 kb, this annotation will be used to check if the pod is already instrumented
	// It will be deleted after mutation
	InstAnnotationInstrumented = domain + "/injected-container"

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
	APMInitContainerNameJava   = "ogas-instrument-otel-java"
	APMInitContainerNameGo     = "ogas-instrument-otel-go"
	APMInitContainerNamePython = "ogas-instrument-otel-python"
	APMInitContainerNameDotNet = "ogas-instrument-otel-dotnet"

	APMVolumeName = "ogas-instrument"
)

// Common Trace Environment Variables
const (
	EnvExporterEndpoint = "OTEL_EXPORTER_OTLP_ENDPOINT"
	EnvTraceSampler     = "OTEL_TRACES_SAMPLER"
	EnvTraceSamplerArg  = "OTEL_TRACES_SAMPLER_ARG"
	EnvPropagators      = "OTEL_PROPAGATORS"
	EnvResourceAttrs    = "OTEL_RESOURCE_ATTRIBUTES"
	EnvServiceName      = "OTEL_SERVICE_NAME"
	EnvTracesExporter   = "OTEL_TRACES_EXPORTER"
	EnvMetricsExporter  = "OTEL_METRICS_EXPORTER"
	EnvLogsExporter     = "OTEL_LOGS_EXPORTER"
)

// Predefined Errors
var (
	ErrEndpointNotDefined = errors.New("endpoint is not defined")
	ErrNotDefined         = errors.New("not defined")
	ErrNotValid           = errors.New("not valid value")
)
