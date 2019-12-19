package kubeflow

import (
	"github.com/CiscoAI/kubeflow-test-infra/pkg/kubeflow/kfcomponent"
	"istio.io/istio/pkg/test/framework/resource"
)

func kfsetup(c resource.Context) error {
	_, err := c.CreateTmpDirectory("kf_test_app")
	if err != nil {
		return err
	}

	kf, err := kfcomponent.New(c, kfcomponent.Config{})
	if err != nil {
		return err
	}
	return nil
}

func setupNative(_ resource.Context) error {
	return nil
}

func setupKube(_ resource.Context) error {
	return nil
}
