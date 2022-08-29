package dependency

import (
	"wire/internal/domain/something/service"
	"wire/internal/domain/something/service/impl"
)

func ProvideService(
	serviceDependency ServiceDependency,
) service.Service {
	return impl.NewService(
		serviceDependency.Repository,
	)
}
