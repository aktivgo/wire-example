//go:build wireinject

package dependency

import (
	"wire/pkg/logging"

	"github.com/google/wire"
)

type CommonDependency struct {
	Logger logging.Logger
}

func GetCommonDependency() CommonDependency {
	panic(
		wire.Build(
			wire.Struct(new(CommonDependency), "*"),

			ProvideLogger,
		),
	)
}

func SingletonGetCommonDependency() CommonDependency {
	return Singleton(
		"GetCommonDependency",
		GetCommonDependency,
	)
}
