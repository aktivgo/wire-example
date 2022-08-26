package dependency

import "github.com/google/wire"

type SomethingRepositoryDependency struct {
	Message string
}

func GetSomethingRepositoryDependency(
	message string,
) SomethingRepositoryDependency {
	panic(
		wire.Build(
			wire.Struct(new(SomethingRepositoryDependency), "*"),

			ProvideMessage,
		),
	)
}
