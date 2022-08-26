package app

import (
	"wire/internal/app/dependency"
)

func Run() {
	server := dependency.GetSomethingHttpServerDependency()
	server.CommonDependency.Logger
}
