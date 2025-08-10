//go:build wireinject
// +build wireinject

//go:generate wire
package console

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	Provide,
	wire.Bind(new(Console), new(*consoleImpl)),
)

func Wire() (Console, error) {
	wire.Build(ProviderSet)
	return &consoleImpl{}, nil
}
