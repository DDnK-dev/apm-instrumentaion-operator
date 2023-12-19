# Development

This document describes how to setup a development environment for the project and initial setup.

## Kubebuilder

The project uses [kubebuilder](https://book.kubebuilder.io/) to generate the boilerplate code for the operator. The
following steps are required to setup the development environment.

### Install kubebuilder

```shell
# download kubebuilder and install locally.
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

### Create a new project

```shell
kubebuilder init --domain ogas.kr --repo github.com/DDnK-dev/apm-instrumentaion-operator --license apache2
kubebuilder create api --version v1 --group apm --kind Instrumentation
kubebuilder create webhook --version v1 --group apm --kind Instrumentation --defaulting --programmatic-validation
```

### Implementation points

- [ ] Implement the `Instrumentation` CRD at `api/v1/instrumentation_types.go`
- [ ] Implement the `Instrumentation` webhook at `api/v1/instrumentation_webhook.go`
- [ ] Implement the `pod` webhook. This webhook will inject the sidecar container into the pod, but not generated from the
      kubebuilder.

