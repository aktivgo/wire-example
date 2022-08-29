package dependency

import (
	"wire/internal/domain/something/handler"
	"wire/internal/domain/something/handler/impl"
)

func ProvideHandler(
	handlerDependency HandlerDependency,
) handler.Handler {
	return impl.NewHandler(
		handlerDependency.Service,
	)
}
