package evaluating

import (
	canarydeploymentmetrics "github.com/solo-io/autopilot/test/e2e/canary/pkg/metrics"
	parameters "github.com/solo-io/autopilot/test/e2e/canary/pkg/parameters"
)

type Inputs struct {
	Metrics         canarydeploymentmetrics.CanaryDeploymentMetrics
	VirtualServices parameters.VirtualServices
}

// FindVirtualService returns <VirtualService, true> if the item is found. else parameters.VirtualService{}, false
func (i Inputs) FindVirtualService(name, namespace string) (parameters.VirtualService, bool) {
	for _, item := range i.VirtualServices.Items {
		if item.Name == name && item.Namespace == namespace {
			return item, true
		}
	}
	return parameters.VirtualService{}, false
}

type Outputs struct {
	VirtualServices parameters.VirtualServices
}
