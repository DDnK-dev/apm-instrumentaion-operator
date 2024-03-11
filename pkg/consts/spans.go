package consts

type ValidationSet map[string]struct{}

func (v ValidationSet) IsInSet(s string) bool {
	if _, ok := v[s]; ok {
		return true
	}
	return false
}

func (v ValidationSet) NotInSet(s string) bool {
	return !v.IsInSet(s)
}

// defines samplers
const (
	AlwaysOnSampler                     = "always_on"
	AlwaysOffSampler                    = "always_off"
	TraceIdRatioBasedSampler            = "traceidratio"
	ParentBasedAlwaysOnSampler          = "parentbased_always_on"
	ParentBasedAlwaysOffSampler         = "parentbased_always_off"
	ParentBasedTraceIdRatioBasedSampler = "parentbased_traceidratio"
	ParentBasedJaegerRemoteSampler      = "parentbased_jaeger_remote"
	JaegerRemoteSampler                 = "jaeger_remote"
	XRaySampler                         = "xray"
)

// SamplerList defines sampler args
var SamplerSet = ValidationSet{
	AlwaysOnSampler:                     {},
	AlwaysOffSampler:                    {},
	TraceIdRatioBasedSampler:            {},
	ParentBasedAlwaysOnSampler:          {},
	ParentBasedAlwaysOffSampler:         {},
	ParentBasedTraceIdRatioBasedSampler: {},
	ParentBasedJaegerRemoteSampler:      {},
	JaegerRemoteSampler:                 {},
	XRaySampler:                         {},
}

// defines propagators
const (
	TraceContextPropagator = "tracecontext"
	BaggagePropagator      = "baggage"
	B3Propagator           = "b3"
	B3MultiPropagator      = "b3multi"
	JaegerPropagator       = "jaeger"
	XRayPropagator         = "xray"
	OTTracePropagator      = "ottrace"
	NonePropagator         = "none"
)

var PropagatorSet = ValidationSet{
	TraceContextPropagator: {},
	BaggagePropagator:      {},
	B3Propagator:           {},
	B3MultiPropagator:      {},
	JaegerPropagator:       {},
	XRayPropagator:         {},
	OTTracePropagator:      {},
	NonePropagator:         {},
}

// defines trace exporters
const (
	OtlpTExporter   = "otlp"
	JaegerTExporter = "jaeger"
	ZipkinTExporter = "zipkin"
	NoneTExporter   = "none"
)

var TraceExporterSet = ValidationSet{
	OtlpTExporter:   {},
	JaegerTExporter: {},
	ZipkinTExporter: {},
	NoneTExporter:   {},
}

// define metrics exporters
const (
	OtlpMExporter = "otlp"
)

//"otlp": OTLP
//"jaeger": export in Jaeger data model
//"zipkin": Zipkin
//"none": No automatically configured exporter for traces.

// defines metrics

//"otlp": OTLP
//"prometheus": Prometheus
//"none": No automatically configured exporter for metrics.
