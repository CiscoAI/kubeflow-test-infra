package argo

import (
	argo "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	log "github.com/sirupsen/logrus"
)

// CreateWorkflow is common function to be used for initializing Argo workflows
func CreateWorkflow() error {
	workflow := argo.Workflow{}
	log.Infof("Workflow intialized: %v", workflow)
	return nil
}
