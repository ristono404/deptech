package container

import (
	"log"

	"github.com/ristono404/deptech/internal/config"
	categoryRepo "github.com/ristono404/deptech/internal/repository/category"
	productRepo "github.com/ristono404/deptech/internal/repository/product"
	userRepo "github.com/ristono404/deptech/internal/repository/user"
	"github.com/ristono404/deptech/internal/shared/database"
	categoryUsecase "github.com/ristono404/deptech/internal/usecase/category"
	productUsecase "github.com/ristono404/deptech/internal/usecase/product"
	userUsecase "github.com/ristono404/deptech/internal/usecase/user"
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
