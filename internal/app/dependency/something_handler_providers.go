package dependency

import (
	"wire/internal/domain/something/repository"
	"wire/internal/domain/something/service"
	"wire/internal/domain/something/service/impl"
)

func ProvideSomethingService(
	somethingRepository repository.SomethingRepository,
) service.SomethingService {
	return impl.NewService(
		somethingRepository,
	)
}
