package utils

import corev1 "k8s.io/api/core/v1"

// LabelMap is pod's labelmap, used to make a plan for mutation
type LabelMap map[string]string

// NewLabelMap returns a new labelmap from the pod
func NewLabelMap(pod *corev1.Pod) LabelMap {
	l := make(map[string]string)
	for k, v := range pod.Labels {
		l[k] = v
	}
	return l
}

// GetLabelValue returns the value of the key from the labelmap, and exists or not
func (l LabelMap) GetLabelValue(key string) (string, bool) {
	a, b := l[key]
	return a, b
}
