package dependency

import (
	"wire/internal/domain/something/handler"

	"github.com/google/wire"
)

type SomethingHttpServerDependency struct {
	CommonDependency           CommonDependency
	SomethingHandlerDependency SomethingHandlerDependency

	SomethingHandler handler.SomethingHandler
}

func GetSomethingHttpServerDependency() SomethingHttpServerDependency {
	panic(
		wire.Build(
			wire.Struct(new(SomethingHttpServerDependency), "*"),

			GetCommonDependency,
			GetSomethingHandlerDependency,

			ProvideSomethingHandler,
		),
	)
}
