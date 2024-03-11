/*
Package resource implements functionalities based on github.com/open-telemetry/opentelemetry-operator/pkg/instrumentation
used to implement the resource inject feature for the operator
*/

package resource

import (
	"context"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	corev1 "k8s.io/api/core/v1"
)

// CreateResourceMap creates resource attribute map.
// User defined attributes (in explicitly set env var) have higher precedence.
func CreateResourceMap(ctx context.Context, otelinst v1alpha1.Instrumentation, ns corev1.Namespace, pod corev1.Pod, index int) map[string]string {
	// get existing resources env var and parse it into a map
	existingRes := map[string]bool{}
	existingResourceEnvIdx := getIndexOfEnv(pod.Spec.Containers[index].Env, constants.EnvOTELResourceAttrs)
	if existingResourceEnvIdx > -1 {
		existingResArr := strings.Split(pod.Spec.Containers[index].Env[existingResourceEnvIdx].Value, ",")
		for _, kv := range existingResArr {
			keyValueArr := strings.Split(strings.TrimSpace(kv), "=")
			if len(keyValueArr) != 2 {
				continue
			}
			existingRes[keyValueArr[0]] = true
		}
	}

	res := map[string]string{}
	for k, v := range otelinst.Spec.Resource.Attributes {
		if !existingRes[k] {
			res[k] = v
		}
	}
	k8sResources := map[attribute.Key]string{}
	k8sResources[semconv.K8SNamespaceNameKey] = ns.Name
	k8sResources[semconv.K8SContainerNameKey] = pod.Spec.Containers[index].Name
	// Some fields might be empty - node name, pod name
	// The pod name might be empty if the pod is created form deployment template
	k8sResources[semconv.K8SPodNameKey] = pod.Name
	k8sResources[semconv.K8SPodUIDKey] = string(pod.UID)
	k8sResources[semconv.K8SNodeNameKey] = pod.Spec.NodeName
	k8sResources[semconv.ServiceInstanceIDKey] = createServiceInstanceId(ns.Name, pod.Name, pod.Spec.Containers[index].Name)
	i.addParentResourceLabels(ctx, otelinst.Spec.Resource.AddK8sUIDAttributes, ns, pod.ObjectMeta, k8sResources)
	for k, v := range k8sResources {
		if !existingRes[string(k)] && v != "" {
			res[string(k)] = v
		}
	}
	return res
}

// creates the service.instance.id following the semantic defined by
// https://github.com/open-telemetry/semantic-conventions/pull/312.
func createServiceInstanceId(namespaceName, podName, containerName string) string {
	var serviceInstanceId string
	if namespaceName != "" && podName != "" && containerName != "" {
		resNames := []string{namespaceName, podName, containerName}
		serviceInstanceId = strings.Join(resNames, ".")
	}
	return serviceInstanceId
}
