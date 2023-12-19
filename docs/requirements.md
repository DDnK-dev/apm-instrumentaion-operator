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

## Instrumentation obejct priority

Which Instrumentation spec should be applied to the application?

- If user didn't set instrumentation,
  - 1. Instrumentation object on the same namespace with application (on multiple case, first one (random) will be applied).
  - 2. Default instrumentation object will be applied, which is labeled with `apm.ogas.kr/default: true`.
- If user set instrumentation, user's instrumentation will be applied.
