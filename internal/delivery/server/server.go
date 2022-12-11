package server

import (
	"github.com/ristono404/deptech/internal/delivery/container"
	httpServer "github.com/ristono404/deptech/internal/delivery/server/http"
)

func New(container *container.Container) {
	httpServer.New(container)
}
