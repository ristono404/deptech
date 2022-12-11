package container

import (
	"log"

	"github.com/deptech/internal/config"
	categoryRepo "github.com/deptech/internal/repository/category"
	productRepo "github.com/deptech/internal/repository/product"
	userRepo "github.com/deptech/internal/repository/user"
	"github.com/deptech/internal/shared/database"
	categoryUsecase "github.com/deptech/internal/usecase/category"
	productUsecase "github.com/deptech/internal/usecase/product"
	userUsecase "github.com/deptech/internal/usecase/user"
)

type Container struct {
	UserService     userUsecase.Service
	CategoryService categoryUsecase.Service
	ProductService  productUsecase.Service
}

func New() *Container {
	c := config.New()

	db, err := database.New(c.Database)
	if err != nil {
		log.Panic(err)
	}

	userRepo := userRepo.New(db)
	categoryRepo := categoryRepo.New(db)
	productRepo := productRepo.New(db)

	userService := userUsecase.New(userRepo)
	categoryService := categoryUsecase.New(categoryRepo)
	productService := productUsecase.New(productRepo)

	return &Container{
		UserService:     userService,
		CategoryService: categoryService,
		ProductService:  productService,
	}
}
