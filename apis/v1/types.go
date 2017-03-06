package v1

import (
	"k8s.io/client-go/1.4/pkg/api/v1"
	autoscaling "k8s.io/client-go/1.4/pkg/apis/autoscaling/v1"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
)

type MemHPAScaler struct {
	v1.ObjectMeta `json:"metadata,omitempty"`
	Spec MemHPASpec `json:"spec,omitempty"`
	Status MemHPAScalerStatus `json:"status,omitempty"`
}

type MemHPASpec struct {
	ScaleTargetRef autoscaling.CrossVersionObjectReference `json:"scaleTargetRef"`
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	MaxReplicas int32 `json:"maxReplicas"`
	TargetUtilizationPercentage *int32 `json:"targetUtilizationPercentage,omitempty"`
}

type MemHPAScalerStatus struct {
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	LastScaleTime *unversioned.Time `json:"lastScaleTime,omitempty"`
	CurrentReplicas int32 `json:"currentReplicas"`
	DesiredReplicas int32 `json:"desiredReplicas"`
	CurrentUtilizationPercentage *int32 `json:"currentCPUUtilizationPercentage,omitempty"`
}