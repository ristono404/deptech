package main

import (
	"github.com/deptech/internal/delivery/container"
	"github.com/deptech/internal/delivery/server"
)

func main() {
	server.New(container.New())
}
