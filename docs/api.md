---
title: "API reference"
description: "APM instrumentation operator generated API reference docs"
---
> This page is automatically generated with `gen-crd-api-reference-docs`.
<p>Packages:</p>
<ul>
<li>
<a href="#apm.ogas.kr%2fv1">apm.ogas.kr/v1</a>
</li>
</ul>
<h2 id="apm.ogas.kr/v1">apm.ogas.kr/v1</h2>
Resource Types:
<ul></ul>
<h3 id="apm.ogas.kr/v1.Configuration">Configuration
</h3>
<p>
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.Go">Go</a>, <a href="#apm.ogas.kr/v1.InstrumentationSpec">InstrumentationSpec</a>, <a href="#apm.ogas.kr/v1.Java">Java</a>)
</p>
<div>
<p>Configuration defines the common configuration for all instrumentation</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type
if all tracer value didn&rsquo;t set, set default value
default=otlp</p>
</td>
</tr>
<tr>
<td>
<code>serviceNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceNameLabel defines the label key used to define the service name
if all value didn&rsquo;t set, set default value
default=app.kubernetes.io/name
this value can be shadowed by OTEL_SERVICE_NAME</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a>
if all Propagator didn&rsquo;t set, set the default value
default={tracecontext, baggage}</p>
</td>
</tr>
<tr>
<td>
<code>envVars</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>envVars defines the environment variables to inject
If there is already an env var with the same name, it will be skipped</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics defines whether to enable metrics
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logs defines whether to enable logs
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apm.ogas.kr/v1.Go">Go
</h3>
<p>
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.InstrumentationSpec">InstrumentationSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endpoint</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Endpoint defines the endpoint to send the data to</p>
</td>
</tr>
<tr>
<td>
<code>sampler</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sampler defines the sampler type, if all samplers didn&rsquo;t set, set default value
default=parentbased_traceidratio
ref:<a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler</a></p>
</td>
</tr>
<tr>
<td>
<code>samplerArg</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SamplerArg defines the sampler argument [0&hellip;1], which is set to target application as env variable
if all sampler args didn&rsquo;t set, sampler type is dtraceidratio or parentbased_traceidratio, set default value
default=&ldquo;0.01&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type
if all tracer value didn&rsquo;t set, set default value
default=otlp</p>
</td>
</tr>
<tr>
<td>
<code>serviceNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceNameLabel defines the label key used to define the service name
if all value didn&rsquo;t set, set default value
default=app.kubernetes.io/name
this value can be shadowed by OTEL_SERVICE_NAME</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a>
if all Propagator didn&rsquo;t set, set the default value
default={tracecontext, baggage}</p>
</td>
</tr>
<tr>
<td>
<code>envVars</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>envVars defines the environment variables to inject
If there is already an env var with the same name, it will be skipped</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics defines whether to enable metrics
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logs defines whether to enable logs
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apm.ogas.kr/v1.Instrumentation">Instrumentation
</h3>
<div>
<p>Instrumentation is the Schema for the instrumentations API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apm.ogas.kr/v1.InstrumentationSpec">
InstrumentationSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>endpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Endpoint defines the endpoint to send the data to</p>
</td>
</tr>
<tr>
<td>
<code>sampler</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sampler defines the sampler type, if all samplers didn&rsquo;t set, set default value
default=parentbased_traceidratio
ref:<a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler</a></p>
</td>
</tr>
<tr>
<td>
<code>samplerArg</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SamplerArg defines the sampler argument [0&hellip;1], which is set to target application as env variable
if all sampler args didn&rsquo;t set, sampler type is dtraceidratio or parentbased_traceidratio, set default value
default=&ldquo;0.01&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type
if all tracer value didn&rsquo;t set, set default value
default=otlp</p>
</td>
</tr>
<tr>
<td>
<code>serviceNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceNameLabel defines the label key used to define the service name
if all value didn&rsquo;t set, set default value
default=app.kubernetes.io/name
this value can be shadowed by OTEL_SERVICE_NAME</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a>
if all Propagator didn&rsquo;t set, set the default value
default={tracecontext, baggage}</p>
</td>
</tr>
<tr>
<td>
<code>envVars</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>envVars defines the environment variables to inject
If there is already an env var with the same name, it will be skipped</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics defines whether to enable metrics
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logs defines whether to enable logs
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>java</code><br/>
<em>
<a href="#apm.ogas.kr/v1.Java">
Java
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>go</code><br/>
<em>
<a href="#apm.ogas.kr/v1.Go">
Go
</a>
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apm.ogas.kr/v1.InstrumentationStatus">
InstrumentationStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apm.ogas.kr/v1.InstrumentationSpec">InstrumentationSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.Instrumentation">Instrumentation</a>)
</p>
<div>
<p>InstrumentationSpec defines the desired state of Instrumentation</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Endpoint defines the endpoint to send the data to</p>
</td>
</tr>
<tr>
<td>
<code>sampler</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sampler defines the sampler type, if all samplers didn&rsquo;t set, set default value
default=parentbased_traceidratio
ref:<a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler</a></p>
</td>
</tr>
<tr>
<td>
<code>samplerArg</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SamplerArg defines the sampler argument [0&hellip;1], which is set to target application as env variable
if all sampler args didn&rsquo;t set, sampler type is dtraceidratio or parentbased_traceidratio, set default value
default=&ldquo;0.01&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type
if all tracer value didn&rsquo;t set, set default value
default=otlp</p>
</td>
</tr>
<tr>
<td>
<code>serviceNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceNameLabel defines the label key used to define the service name
if all value didn&rsquo;t set, set default value
default=app.kubernetes.io/name
this value can be shadowed by OTEL_SERVICE_NAME</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a>
if all Propagator didn&rsquo;t set, set the default value
default={tracecontext, baggage}</p>
</td>
</tr>
<tr>
<td>
<code>envVars</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>envVars defines the environment variables to inject
If there is already an env var with the same name, it will be skipped</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics defines whether to enable metrics
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logs defines whether to enable logs
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>java</code><br/>
<em>
<a href="#apm.ogas.kr/v1.Java">
Java
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>go</code><br/>
<em>
<a href="#apm.ogas.kr/v1.Go">
Go
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apm.ogas.kr/v1.InstrumentationStatus">InstrumentationStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.Instrumentation">Instrumentation</a>)
</p>
<div>
<p>InstrumentationStatus defines the observed state of Instrumentation</p>
</div>
<h3 id="apm.ogas.kr/v1.Java">Java
</h3>
<p>
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.InstrumentationSpec">InstrumentationSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Image is a container image with javaagent auto-instrumentation JAR.</p>
</td>
</tr>
<tr>
<td>
<code>endpoint</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Endpoint defines the endpoint to send the data to</p>
</td>
</tr>
<tr>
<td>
<code>sampler</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sampler defines the sampler type, if all samplers didn&rsquo;t set, set default value
default=parentbased_traceidratio
ref:<a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler</a></p>
</td>
</tr>
<tr>
<td>
<code>samplerArg</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SamplerArg defines the sampler argument [0&hellip;1], which is set to target application as env variable
if all sampler args didn&rsquo;t set, sampler type is dtraceidratio or parentbased_traceidratio, set default value
default=&ldquo;0.01&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type
if all tracer value didn&rsquo;t set, set default value
default=otlp</p>
</td>
</tr>
<tr>
<td>
<code>serviceNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceNameLabel defines the label key used to define the service name
if all value didn&rsquo;t set, set default value
default=app.kubernetes.io/name
this value can be shadowed by OTEL_SERVICE_NAME</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a>
if all Propagator didn&rsquo;t set, set the default value
default={tracecontext, baggage}</p>
</td>
</tr>
<tr>
<td>
<code>envVars</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>envVars defines the environment variables to inject
If there is already an env var with the same name, it will be skipped</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics defines whether to enable metrics
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logs defines whether to enable logs
if all value didn&rsquo;t set, set default value
default=none</p>
</td>
</tr>
<tr>
<td>
<code>logging</code><br/>
<em>
string
</em>
</td>
<td>
<p>Logging defines the logging configuration
kubebuilder:default=simple
kubebuilder:validation:Enum=simple;none;application</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apm.ogas.kr/v1.Sampling">Sampling
</h3>
<p>
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.Go">Go</a>, <a href="#apm.ogas.kr/v1.InstrumentationSpec">InstrumentationSpec</a>, <a href="#apm.ogas.kr/v1.Java">Java</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sampler</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sampler defines the sampler type, if all samplers didn&rsquo;t set, set default value
default=parentbased_traceidratio
ref:<a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#sampler</a></p>
</td>
</tr>
<tr>
<td>
<code>samplerArg</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SamplerArg defines the sampler argument [0&hellip;1], which is set to target application as env variable
if all sampler args didn&rsquo;t set, sampler type is dtraceidratio or parentbased_traceidratio, set default value
default=&ldquo;0.01&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<hr/>
