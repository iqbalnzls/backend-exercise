package main

import (
	"github.com/iqbalnzls/backend-exercise/internal/delivery/rest"
	"github.com/iqbalnzls/backend-exercise/internal/di"
)

func main() {
	rest.StartHttpServer(di.SetupContainer())
}
