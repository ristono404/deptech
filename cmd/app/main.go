package main

import (
	"github.com/ristono404/deptech/internal/delivery/container"
	"github.com/ristono404/deptech/internal/delivery/server"
)

func main() {
	server.New(container.New())
}
