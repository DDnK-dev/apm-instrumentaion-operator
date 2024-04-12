package utils

import corev1 "k8s.io/api/core/v1"

// AnnotationMap is pod's labelmap, used to make a plan for mutation
type AnnotationMap map[string]string

// NewAnnotationMap returns a new labelmap from the pod
func NewAnnotationMap(pod *corev1.Pod) AnnotationMap {
	l := make(map[string]string)
	for k, v := range pod.Annotations {
		l[k] = v
	}
	return l
}

// GetAnnotationValue returns the value of the key from the labelmap, and exists or not
func (l AnnotationMap) GetAnnotationValue(key string) (string, bool) {
	a, b := l[key]
	return a, b
}

func In(a string, b ...string) bool {
	for _, c := range b {
		if a == c {
			return true
		}
	}
	return false
}
