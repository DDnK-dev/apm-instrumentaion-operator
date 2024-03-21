# Annotations

This document describes the annotations that can be used to configure instrumentation on the application.

## Annotations List

All annotations are set to `spec.template.metadata.annotations`

### Java

| annotation                      | description                                                                                                                            | required |
|---------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|:---------|
| apm.ogas.kr/inject-java         | Set Instrumentation used to apm injection                                                                                              | O        |
| apm.ogas.kr/java-containe-names | Set the container names to inject. If this annotations is not set on multiple container pod, only first container will be instrumented | X        |

#### `apm.ogas.kr/inject-java`

Possible values:
- `true` : Get the instrumentation resource from same namespace. If not exists, or multiple resources exist, inject will be failed.
- `false` : Skip the injection.
- `my-ns/my-instrumentation` : Get the instrumentation resource named `my-instrumentation` from `my-ns` namespace. If not exists inject will be failed.

#### Example

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-pod
  namespace: my-ns
spec:
  template:
    metadata:
    annotations:
      apm.ogas.kr/inject-java: "true"
```
