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
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/consts"
	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
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
	if r.Spec.Sampling.Sampler == "" {
		r.Spec.Sampling.Sampler = "parentbased_traceidratio"
	}
	if r.Spec.Sampling.SamplerArg == "" && utils.In(r.Spec.Sampling.Sampler,
		consts.TraceIdRatioBasedSampler, consts.ParentBasedTraceIdRatioBasedSampler) {
		r.Spec.Sampling.SamplerArg = "0.1"
	}
	if r.Spec.Configuration.Tracer == "" {
		r.Spec.Configuration.Tracer = "otlp"
	}
	if r.Spec.Configuration.ServiceNameLabel == "" {
		r.Spec.Configuration.ServiceNameLabel = "app.kubernetes.io/name"
	}
	if len(r.Spec.Configuration.Propagator) == 0 {
		r.Spec.Configuration.Propagator = []string{consts.B3Propagator, consts.JaegerPropagator}
	}
	if r.Spec.Configuration.Metrics == "" {
		r.Spec.Configuration.Metrics = "none"
	}
	if r.Spec.Configuration.Logs == "" {
		r.Spec.Configuration.Logs = "none"
	}
}

//+kubebuilder:webhook:path=/validate-apm-ogas-kr-v1-instrumentation,mutating=false,failurePolicy=fail,sideEffects=None,groups=apm.ogas.kr,resources=instrumentations,verbs=create;update,versions=v1,name=vinstrumentation.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Instrumentation{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Instrumentation) ValidateCreate() (admission.Warnings, error) {
	instrumentationlog.Info("validate create", "name", r.Name)
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Instrumentation) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	instrumentationlog.Info("validate update", "name", r.Name)
	return nil, nil
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
