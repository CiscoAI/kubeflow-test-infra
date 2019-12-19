package kfcomponent

import (
	"istio.io/istio/pkg/test/framework/components/environment/kube"
	"istio.io/istio/pkg/test/framework/resource"
)

type kubeComponent struct {
	id resource.ID
}

func newKube(ctx resource.Context, _ Config) Instance {
	i := &kubeComponent{}

	ctx.TrackResource(i)

	env := ctx.Environment().(*kube.Environment)
	_ = env

	return i
}
