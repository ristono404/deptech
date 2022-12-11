package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ristono404/deptech/internal/delivery/server/http/handler"
	"github.com/ristono404/deptech/internal/pkg/midleware"
)

func New(r *mux.Router, handler *handler.Handler) {

	r.HandleFunc("/login", handler.User.Login)

	rLogout := r.PathPrefix("").Subrouter()
	rLogout.HandleFunc("/logout", handler.User.Logout)
	rLogout.Use(midleware.AuthUser)

	v0Auth := r.PathPrefix("/v0").Subrouter()
	v0Auth.HandleFunc("/users", handler.User.List).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.HandleFunc("/user/{id}", handler.User.Read).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.HandleFunc("/users", handler.User.Create).Methods(http.MethodPost, http.MethodOptions)
	v0Auth.HandleFunc("/user/{id}", handler.User.Update).Methods(http.MethodPut, http.MethodOptions)
	v0Auth.HandleFunc("/users", handler.User.Delete).Methods(http.MethodDelete, http.MethodOptions)

	v0Auth.HandleFunc("/categories", handler.Category.List).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.HandleFunc("/category/{id}", handler.Category.Read).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.HandleFunc("/categories", handler.Category.Create).Methods(http.MethodPost, http.MethodOptions)
	v0Auth.HandleFunc("/category/{id}", handler.Category.Update).Methods(http.MethodPut, http.MethodOptions)
	v0Auth.HandleFunc("/categories", handler.Category.Delete).Methods(http.MethodDelete, http.MethodOptions)

	v0Auth.HandleFunc("/products", handler.Product.List).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.HandleFunc("/product/{id}", handler.Product.Read).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.HandleFunc("/products", handler.Product.Create).Methods(http.MethodPost, http.MethodOptions)
	v0Auth.HandleFunc("/product/{id}", handler.Product.Update).Methods(http.MethodPut, http.MethodOptions)
	v0Auth.HandleFunc("/products", handler.Product.Delete).Methods(http.MethodDelete, http.MethodOptions)

	v0Auth.HandleFunc("/transaction/in", handler.Product.In).Methods(http.MethodPost, http.MethodOptions)
	v0Auth.HandleFunc("/transaction/out", handler.Product.Out).Methods(http.MethodPost, http.MethodOptions)
	v0Auth.HandleFunc("/transaction", handler.Product.TransactionHistory).Methods(http.MethodGet, http.MethodOptions)
	v0Auth.Use(midleware.AuthUser)

}
