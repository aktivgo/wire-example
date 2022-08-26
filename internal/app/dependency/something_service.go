package dependency

import "github.com/google/wire"

type SomethingServiceDependency struct {
	SomethingRepositoryDependency SomethingHandlerDependency
}

func GetSomethingServiceDepenedency() SomethingServiceDependency {
	panic(
		wire.Build(
			wire.Struct(new(SomethingServiceDependency), "*"),

			GetSomethingRepositoryDependency,
		),
	)
}
