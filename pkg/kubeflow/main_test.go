package kubeflow

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/environment"
	"istio.io/istio/pkg/test/framework/components/environment/kube"
	"istio.io/istio/pkg/test/framework/components/istio"
)

var (
	ist istio.Instance
	env *kube.Environment
)

// TestMain for kfctl test suite
func TestMain(m *testing.M) {
	framework.
		NewSuite("kfctl_test", m).
		RequireEnvironment(environment.Kube).
		SetupOnEnv(environment.Kube, istio.Setup(&i, setupKubeflow)).
		Run()
}

func setupKubeflow(cfg *istio.Config) {
	log.Infof("Config environment, value: %v", cfg.Values["istio.test.env"])
}

func TestSmoke(t *testing.T) {
	// Get test context
	ctx := framework.GetContext(t)
	defer ctx.Done()
}
