package kfcomponent

import (
	kfdefsv1beta1 "github.com/kubeflow/kfctl/v3/pkg/apis/apps/kfdef/v1beta1"
	"istio.io/istio/pkg/test/framework/components/environment"
	"istio.io/istio/pkg/test/framework/resource"
)

type Instance iterface {
	resource.Resource
}

type Config struct {
	KfDef *kfdefsv1beta1.KfDef
}

func New(ctx resource.Context, cfg Config) (i Instance, err error) {
	switch ctx.Environment().EnvironmentName() {
	case environment.Native:
		i = newNative(ctx, cfg)
	case environment.Kube:
		i = newKube(ctx, cfg)
	default:
		err = resource.UnsupportedEnvironment(ctx.Environment())
	}
	return i, err
}

func NewOrFail(t *testing.T, ctx resource.Context, cfg Config) Instance {
	t.Helper()

	i, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("kfcomponent.NeworFail: %v", err)
	}
	return i
}
