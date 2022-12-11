package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ristono404/deptech/internal/delivery/container"
	"github.com/ristono404/deptech/internal/delivery/server/http/handler"
	"github.com/ristono404/deptech/internal/delivery/server/http/handler/router"
)

func New(container *container.Container) {
	r := mux.NewRouter()
	router.New(r, handler.New(container))
	fmt.Println("⇨ http server started on \033[32m[::]:8080\033[0m")
	log.Println(http.ListenAndServe(":8080", handlers.CompressHandler(r)))
}
