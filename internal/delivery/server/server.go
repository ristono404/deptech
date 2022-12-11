package server

import (
	"github.com/deptech/internal/delivery/container"
	httpServer "github.com/deptech/internal/delivery/server/http"
)

func New(container *container.Container) {
	httpServer.New(container)
}
