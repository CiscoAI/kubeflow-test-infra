package kfcomponent

import (
	"istio.io/istio/pkg/test/framework/components/environment/native"
	"istio.io/istio/pkg/test/framework/resource"
)

type nativeComponent struct {
	is resource.ID
}

func newNative(ctx resource.Context, _ Config) Instance {
	i := &nativeComponent{}

	ctx.TrackResource(i)

	env := ctx.environment().(*native.Environment)
	_ = env

	return i
}
