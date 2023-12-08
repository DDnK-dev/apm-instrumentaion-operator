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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InstrumentationSpec defines the desired state of Instrumentation
type InstrumentationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Endpoint defines the endpoint to send the data to
	// +kubebuilder:required
	Endpoint string `json:"endpoint"`

	// Configuration defines the common configuration for all instrumentation
	// root configuration may be overridden by each instrumentation
	Configuration `json:",inline"`

	Java Java `json:"java,omitempty"`

	Go Go `json:"go,omitempty"`
}

// Configuration defines the common configuration for all instrumentation
type Configuration struct {
	// Tracer defines the tracer type
	// +kubebuilder:default=otlp
	// +kubebuilder:validation:Enum=none;otlp;jaeger;zipkin;logging
	// +optional
	Tracer string `json:"tracer,omitempty"`

	// ServiceNameLabel defines the label key used to define the service name
	// +kubebuilder:default=app.kubernetes.io/name
	ServiceNameLabel string `json:"serviceNameLabel,omitempty"`

	// Propagator defines the propagation type, comma-separated list of propagators
	// ref: https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator
	// +kubebuilder:default=tracecontext,baggage
	Propagator string `json:"propagator,omitempty"`

	// Sampler defines the sampler type
	// ref:https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler
	// +kubebuilder:default=parentbased_traceidratio
	Sampler string `json:"sampler,omitempty"`

	// SamplerArg defines the sampler argument [0...1]
	// +kubebuilder:default=0.01
	SamplerArg string `json:"samplerArg,omitempty"`

	// Metrics defines whether to enable metrics
	// +kubebuilder:default=none
	// +kubebuilder:validation:Enum=none;otlp;logging;prometheus
	Metrics string `json:"metrics,omitempty"`

	// Logs defines whether to enable logs
	// +kubebuilder:default=none
	// +kubebuilder:validation:Enum=none;otlp;logging
	Logs string `json:"logs,omitempty"`
}

type Java struct {
	// Endpoint defines the endpoint to send the data to
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// Configuration defines the common configuration for all instrumentation
	Configuration `json:",inline"`

	// Logging defines the logging configuration
	// kubebuilder:default=simple
	// kubebuilder:validation:Enum=simple;none;application
	Logging string `json:"logging,omitempty"`
}

type Go struct {
	// Endpoint defines the endpoint to send the data to
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// GoTarget defines the executable target to instrument
	// +kubebuilder:required
	GoTarget string `json:"goTarget"`

	// Configuration defines the common configuration for all instrumentation
	Configuration `json:",inline"`
}

// InstrumentationStatus defines the observed state of Instrumentation
type InstrumentationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Instrumentation is the Schema for the instrumentations API
type Instrumentation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstrumentationSpec   `json:"spec,omitempty"`
	Status InstrumentationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InstrumentationList contains a list of Instrumentation
type InstrumentationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instrumentation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Instrumentation{}, &InstrumentationList{})
}
