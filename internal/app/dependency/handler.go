//go:build wireinject

package dependency

import (
	"wire/internal/domain/something/service"

	"github.com/google/wire"
)

type HandlerDependency struct {
	ServiceDependency ServiceDependency

	Service service.Service
}

func GetHandlerDependency() HandlerDependency {
	panic(
		wire.Build(
			wire.Struct(new(HandlerDependency), "*"),

			GetServiceDepenedency,

			ProvideService,
		),
	)
}
