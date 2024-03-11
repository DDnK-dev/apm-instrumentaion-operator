package utils

import corev1 "k8s.io/api/core/v1"

func HasInitContainer(pod *corev1.Pod, name string) bool {
	for _, cont := range pod.Spec.InitContainers {
		if cont.Name == name {
			return true
		}
	}
	return false
}

func HasAnnotation(annotations map[string]string, key string) bool {
	if _, ok := annotations[key]; ok {
		return true
	}
	return false
}

func GetContIndexWithName(name string, pod *corev1.Pod) int {
	for i, cont := range pod.Spec.Containers {
		if cont.Name == name {
			return i
		}
	}
	return -1
}

// NotIn returns true if a is not in b
// hypothesis: b is not long enough to be a set
func NotIn(a string, b ...string) bool {
	for _, c := range b {
		if a == c {
			return false
		}
	}
	return true
}

func In(a string, b ...string) bool {
	for _, c := range b {
		if a == c {
			return true
		}
	}
	return false
}

// NotInSet return true if a is not in set b
func NotInSet(a string, b map[string]bool) bool {
	if _, ok := b[a]; ok {
		return false
	}
	return true
}
