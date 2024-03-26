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

func In(a string, b ...string) bool {
	for _, c := range b {
		if a == c {
			return true
		}
	}
	return false
}
