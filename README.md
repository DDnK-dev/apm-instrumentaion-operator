# apm-instrument-operator

Personal project for studying auto instrumentation of apm and kubernetes operator.

Clone Coding of [Opentelemetry-operator](https://github.com/open-telemetry/opentelemetry-operator),
but auto instrumentation only

This operator is for
- opentelemetry auto instrumentation
- jaeger auto instrumentation (just for some cases like kafka clients)

## Feature

---

### Auto Instrumentation Target

| Type          | Instrumentation |
|:--------------|:----------------|
| opentelemetry | Java            |
| opentelemetry | Python          |
| opentelemetry | Golang          |
| opentelemetry | NodeJS          |
| opentelemetry | Nginx           |
| jaeger        | Java            |

