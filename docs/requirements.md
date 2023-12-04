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

kubebuilder init --domain ogas.kr --repo github.com/DDnK-dev/apm-instrumentaion-operator --license apache2

kubebuilder create api --version v1 --group apm --kind Instrumentation

kubebuilder create webhook --version v1 --group apm --kind Instrumentation --defaulting --programmatic-validation

