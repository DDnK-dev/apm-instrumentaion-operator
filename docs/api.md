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
(<em>Appears on:</em><a href="#apm.ogas.kr/v1.InstrumentationSpec">InstrumentationSpec</a>, <a href="#apm.ogas.kr/v1.Java">Java</a>)
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
<p>Tracer defines the tracer type</p>
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
<p>ServiceNameLabel defines the label key used to define the service name</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
string
</em>
</td>
<td>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a></p>
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
<p>Sampler defines the sampler type
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
<p>SamplerArg defines the sampler argument [0&hellip;1]</p>
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
<p>Metrics defines whether to enable metrics</p>
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
<p>Logs defines whether to enable logs</p>
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
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type</p>
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
<p>ServiceNameLabel defines the label key used to define the service name</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
string
</em>
</td>
<td>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a></p>
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
<p>Sampler defines the sampler type
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
<p>SamplerArg defines the sampler argument [0&hellip;1]</p>
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
<p>Metrics defines whether to enable metrics</p>
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
<p>Logs defines whether to enable logs</p>
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
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type</p>
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
<p>ServiceNameLabel defines the label key used to define the service name</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
string
</em>
</td>
<td>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a></p>
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
<p>Sampler defines the sampler type
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
<p>SamplerArg defines the sampler argument [0&hellip;1]</p>
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
<p>Metrics defines whether to enable metrics</p>
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
<p>Logs defines whether to enable logs</p>
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
<code>tracer</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tracer defines the tracer type</p>
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
<p>ServiceNameLabel defines the label key used to define the service name</p>
</td>
</tr>
<tr>
<td>
<code>propagator</code><br/>
<em>
string
</em>
</td>
<td>
<p>Propagator defines the propagation type, comma-separated list of propagators
ref: <a href="https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator">https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md#propagator</a></p>
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
<p>Sampler defines the sampler type
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
<p>SamplerArg defines the sampler argument [0&hellip;1]</p>
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
<p>Metrics defines whether to enable metrics</p>
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
<p>Logs defines whether to enable logs</p>
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
<hr/>
