# Requirements

- User can instrument their application with opentelemetry and jaeger
- User can instrument their application just by adding some labels to their application
- User can declare their instrumentation spec with CRD
- Operator must validate user's instrumentation spec
- Operator must create a Custom Resource for each instrumentation spec
- User can declare default instrumentation spec for whole namespace
- Operator must create a Custom Resource for each default instrumentation spec

## required modules for operator

- admission controller
- api resource
