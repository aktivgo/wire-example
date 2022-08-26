package dependency

import (
	"wire/internal/domain/something/handler"
	"wire/internal/domain/something/handler/impl"
)

func ProvideSomethingHandler(
	somethingHandlerDependency SomethingHandlerDependency,
) handler.SomethingHandler {
	return impl.NewHandler(
		somethingHandlerDependency.SomethingService,
	)
}
