//go:build wireinject
// +build wireinject

package main

import (
	adapter_in "github.com/Terence1105/hexagonal/account/adapter/in"
	adapter_out "github.com/Terence1105/hexagonal/account/adapter/out"
	"github.com/Terence1105/hexagonal/account/application/service"
	"github.com/google/wire"
)

func InitializeGinAdapter() adapter_in.GinAdapter {
	wire.Build(
		adapter_out.NewAccountLevelDBAdapter,
		service.NewTransferService,
		adapter_in.NewGinAdapter)

	return adapter_in.GinAdapter{}
}

func InitializeFiberAdapter() adapter_in.FiberAdapter {
	wire.Build(
		adapter_out.NewAccountLevelDBAdapter,
		service.NewTransferService,
		adapter_in.NewFiberAdapter)

	return adapter_in.FiberAdapter{}
}
