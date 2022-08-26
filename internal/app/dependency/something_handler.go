package dependency

import (
	"wire/internal/domain/something/service"

	"github.com/google/wire"
)

type SomethingHandlerDependency struct {
	SomethingService service.SomethingService
}

func GetSomethingHandlerDependency() SomethingHandlerDependency {
	panic(
		wire.Build(
			wire.Struct(new(SomethingHandlerDependency), "*"),

			ProvideSomethingService,
		),
	)
}
