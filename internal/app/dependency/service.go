//go:build wireinject

package dependency

import (
	"wire/internal/domain/something/repository"

	"github.com/google/wire"
)

type ServiceDependency struct {
	RepositoryDependency RepositoryDependency

	Repository repository.Repository
}

func GetServiceDepenedency() ServiceDependency {
	panic(
		wire.Build(
			wire.Struct(new(ServiceDependency), "*"),

			GetRepositoryDependency,

			ProvideRepository,
		),
	)
}
