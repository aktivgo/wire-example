//go:build wireinject

package dependency

import "github.com/google/wire"

type RepositoryDependency struct {
	Message string
}

func GetRepositoryDependency() RepositoryDependency {
	panic(
		wire.Build(
			wire.Struct(new(RepositoryDependency), "*"),

			ProvideMessage,
		),
	)
}
