package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ResourceInfo struct {
	Total       string `json:"total"`
	Allocatable string `json:"allocatable"`
	Used        string `json:"used"`
}

type NodeResourceInfoSpec struct {
	NodeName         string                  `json:"nodeName"`
	Resources        map[string]ResourceInfo `json:"resources"`
	Status           string                  `json:"status"`
	Roles            string                  `json:"roles"`
	ScheduleVersion  string                  `json:"scheduleVersion"`
	InternalIp       string                  `json:"internalIp"`
	OS               string                  `json:"os"`
	KernelVersion    string                  `json:"kernelVersion"`
	ContainerRuntime string                  `json:"containerRuntime"`
}

type NodeResourceInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NodeResourceInfoSpec `json:"spec"`
}
