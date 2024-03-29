/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

// log is for logging in this package.
var instrumentationlog = logf.Log.WithName("instrumentation-resource")

// SetupWebhookWithManager will set up the manager to manage the webhooks
func (r *Instrumentation) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-apm-ogas-kr-v1-instrumentation,mutating=true,failurePolicy=fail,sideEffects=None,groups=apm.ogas.kr,resources=instrumentations,verbs=create;update,versions=v1,name=minstrumentation.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Instrumentation{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Instrumentation) Default() {
	instrumentationlog.Info("default", "name", r.Name)
	r.Spec.Sampling.defaulter()
	r.Spec.Configuration.defaulter()
	r.Spec.Java.defaulter()
}

//+kubebuilder:webhook:path=/validate-apm-ogas-kr-v1-instrumentation,mutating=false,failurePolicy=fail,sideEffects=None,groups=apm.ogas.kr,resources=instrumentations,verbs=create;update,versions=v1,name=vinstrumentation.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Instrumentation{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Instrumentation) ValidateCreate() (admission.Warnings, error) {
	instrumentationlog.Info("validate create", "name", r.Name)
	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Instrumentation) ValidateUpdate(_ runtime.Object) (admission.Warnings, error) {
	instrumentationlog.Info("validate update", "name", r.Name)
	return r.validate()
}

// validate validates if Instrumentation object fulfill
func (r *Instrumentation) validate() (admission.Warnings, error) {
	warning := admission.Warnings{}
	if r.Spec.Endpoint == "" {
		return warning, errors.New("spec.endpoint must be specified")
	}
	if r.Spec.Sampling.validate() != nil {
		warning = append(warning, "spec.sampling must be valid")
		return warning, errors.New("spec.sampling must be valid")
	}
	if r.Spec.Configuration.validate() != nil {
		warning = append(warning, "spec.configuration must be valid")
		return warning, errors.New("spec.configuration must be valid")
	}
	if r.Spec.Java.validate() != nil {
		warning = append(warning, "spec.java must be valid")
		return warning, errors.New("spec.java must be valid")
	}
	return warning, nil
}

func (r *Instrumentation) ValidateObject() error {
	if r.Spec.Endpoint == "" {
		return errors.Wrap(consts.ErrNotDefined, "endpoint")
	}
	// sampler validation
	if r.Spec.Sampling.Sampler == "" {
		return errors.Wrap(consts.ErrNotDefined, "sampler")
	} else if consts.SamplerSet.NotInSet(r.Spec.Sampling.Sampler) {
		return errors.Wrap(consts.ErrNotValid, "sampler")
	} else if utils.In(r.Spec.Sampling.Sampler,
		consts.TraceIdRatioBasedSampler, consts.ParentBasedTraceIdRatioBasedSampler) &&
		r.Spec.Sampling.SamplerArg == "" {
		return errors.Wrap(consts.ErrNotDefined, "samplerArg")
	}

	if len(r.Spec.Propagator) == 0 {
		return errors.Wrap(consts.ErrNotDefined, "propagator")
	} else {
		for _, prop := range r.Spec.Propagator {
			if consts.PropagatorSet.NotInSet(prop) {
				return errors.Wrap(consts.ErrNotValid, "propagator: "+prop)
			}
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Instrumentation) ValidateDelete() (admission.Warnings, error) { return nil, nil }

func (s *Sampling) defaulter() {
	if s.Sampler == "" {
		s.Sampler = consts.ParentBasedTraceIdRatioBasedSampler
	}
	if consts.RateSamplerSet.IsInSet(s.Sampler) && s.SamplerArg == "" {
		s.SamplerArg = consts.DefaultSamplingRate
	}
}

func (s *Sampling) validate() error {
	if s.Sampler != "" && consts.SamplerSet.NotInSet(s.Sampler) {
		return errors.Wrap(consts.ErrNotValid, "sampler")
	}
	if s.SamplerArg != "" {
		if _, err := strconv.ParseFloat(s.SamplerArg, 32); err != nil {
			return fmt.Errorf("samplerArg must be a number, got %s", s.SamplerArg)
		}
	}
	return nil
}

func (c *Configuration) defaulter() {
	if c.Tracer == "" {
		c.Tracer = consts.OtlpTExporter
	}
	if c.ServiceNameLabel == "" {
		c.ServiceNameLabel = consts.DefaultServiceNameLabel
	}
	if len(c.Propagator) == 0 {
		c.Propagator = []string{consts.TraceContextPropagator, consts.BaggagePropagator}
	}
	if c.Metrics == "" {
		c.Metrics = consts.NoneMExporter
	}
	if c.Logs == "" {
		c.Logs = consts.NoneLExporter
	}
}

func (c *Configuration) validate() error {
	if c.Tracer != "" && consts.TraceExporterSet.NotInSet(c.Tracer) {
		return errors.Wrap(consts.ErrNotValid, "tracer")
	}
	for _, prop := range c.Propagator {
		if prop != "" && consts.PropagatorSet.NotInSet(prop) {
			return errors.Wrap(consts.ErrNotValid, "propagator: "+prop)
		}
	}
	if c.Metrics != "" && consts.MetricExporterSet.NotInSet(c.Metrics) {
		return errors.Wrap(consts.ErrNotValid, "metrics")
	}
	if c.Logs != "" && consts.LogExporterSet.NotInSet(c.Logs) {
		return errors.Wrap(consts.ErrNotValid, "logs")
	}
	return nil
}

func (j *Java) defaulter() {
	if j.Image == "" {
		j.Image = consts.APMImageJava
	}
	if j.Logging == "" {
		j.Logging = consts.LoggingSimple
	}
}

func (j *Java) validate() error {
	if j.Sampling.Sampler != "" || j.Sampling.SamplerArg != "" {
		if err := j.Sampling.validate(); err != nil {
			return err
		}
	}
	if err := j.Config.validate(); err != nil {
		return err
	}
	if j.Logging != "" && consts.JavaLoggingSet.NotInSet(j.Logging) {
		return errors.Wrap(consts.ErrNotValid, "java-logging")
	}
	return nil
}
