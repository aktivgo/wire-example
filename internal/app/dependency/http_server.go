//go:build wireinject

package dependency

import (
	"wire/internal/domain/something/handler"

	"github.com/google/wire"
)

type HttpServerDependency struct {
	CommonDependency CommonDependency

	HandlerDependency HandlerDependency

	Handler handler.Handler
}

func GetHttpServerDependency() HttpServerDependency {
	panic(
		wire.Build(
			wire.Struct(new(HttpServerDependency), "*"),

			SingletonGetCommonDependency,

			GetHandlerDependency,

			ProvideHandler,
		),
	)
}
