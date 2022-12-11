package handler

import (
	"github.com/deptech/internal/delivery/container"
	"github.com/deptech/internal/delivery/server/http/handler/category"
	"github.com/deptech/internal/delivery/server/http/handler/product"
	"github.com/deptech/internal/delivery/server/http/handler/user"
)

type Handler struct {
	User     *user.Handler
	Category *category.Handler
	Product  *product.Handler
}

func New(container *container.Container) *Handler {
	user := user.New(container.UserService)
	category := category.New(container.CategoryService)
	product := product.New(container.ProductService)

	return &Handler{user, category, product}
}
