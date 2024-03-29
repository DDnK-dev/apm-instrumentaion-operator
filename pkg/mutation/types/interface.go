package types

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/DDnK-dev/apm-instrumentaion-operator/pkg/utils"
)

type Injector interface {
	PlanMutation(*corev1.Pod, utils.AnnotationMap) *Plan // 초기화 로직
	SetClient(client.Client) Injector                    // 클라이언트 설정 로직
	SetPlan(*Plan)                                       // Plan을 설정하는 로직
	GetPlan() *Plan
	Mutate(context.Context, *corev1.Pod) (*corev1.Pod, error) // 실제 변이 로직
}
