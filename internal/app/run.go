package app

import (
	"log"
	"wire/internal/app/dependency"
)

func Run() {
	server := dependency.GetSomethingHttpServerDependency()

	if err := server.SomethingHandler.Handle(); err != nil {
		log.Println(err)
	}
}
