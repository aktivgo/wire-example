package dependency

import (
	"wire/internal/domain/something/repository"
	"wire/internal/domain/something/repository/impl"
)

func ProvideRepository(
	repositoryDependency RepositoryDependency,
) repository.Repository {
	return impl.NewRepository(
		repositoryDependency.Message,
	)
}
