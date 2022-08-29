package app

import (
	"log"
	"wire/internal/app/dependency"
)

func Run() {
	server := dependency.GetHttpServerDependency()

	if err := server.Handler.Handle(); err != nil {
		log.Println(err)
	}
}
