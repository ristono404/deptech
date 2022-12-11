package category

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	config "github.com/ristono404/deptech/internal/config"
	event "github.com/ristono404/deptech/internal/event"
	pkgRequest "github.com/ristono404/deptech/internal/pkg/request"
	response "github.com/ristono404/deptech/internal/pkg/response"
	categoryUsecase "github.com/ristono404/deptech/internal/usecase/category"
)

type Handler struct {
	Service categoryUsecase.Service
	Config  config.Config
}

func New(service categoryUsecase.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	page, limit, offset, from, to, err := pkgRequest.List(r.URL.Query(), h.Config.PerPage)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to get category",
		}, http.StatusInternalServerError)

		return
	}
	data, total, pages, err := h.Service.List(
		int(offset),
		int(limit),
		from,
		to,
		r.URL.Query().Get("search"),
		strings.Split(r.URL.Query().Get("by"), ","),
	)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to get category",
		}, http.StatusInternalServerError)

		return
	}

	if len(data) == 0 {
		response.Response(w, response.View{
			Message: "data not found",
		}, http.StatusOK)

		return
	}

	response.Response(w, response.View{
		Status:  true,
		Data:    data,
		Message: "data found",
		Pagination: response.Pages{
			Total:      total,
			PerPage:    limit,
			Current:    page,
			TotalPages: pages,
		},
	}, http.StatusOK)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var body event.CategoryCreated
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		}, http.StatusBadRequest)
		return
	}

	if err := validator.New().Struct(body); err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "Bad Request",
		}, http.StatusBadRequest)
		return
	}

	data, err := h.Service.Create(
		body,
	)

	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	response.Response(w, response.View{
		Status: true,
		Data:   data,
	}, http.StatusOK)

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	var body event.CategoryCreated
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		}, http.StatusBadRequest)
		return
	}

	body.ID, err = strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "Bad Request",
		}, http.StatusBadRequest)
		return
	}

	if err := validator.New().Struct(body); err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "Bad Request",
		}, http.StatusBadRequest)
		return
	}

	data, err := h.Service.Update(
		body,
	)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	response.Response(w, response.View{
		Status: true,
		Data:   data,
	}, http.StatusOK)

}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Message: "failed to get category",
		}, http.StatusInternalServerError)
		return
	}

	data, err := h.Service.Read(id)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to read the category",
		}, http.StatusInternalServerError)

		return
	}

	if data == nil {
		response.Response(w, response.View{
			Message: "data not found",
		}, http.StatusOK)
		return
	}

	response.Response(w, response.View{
		Status: true,
		Data:   data,
	}, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Ids []uint64 `json:"ids" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println(err)
		response.Response(w, response.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		}, http.StatusBadRequest)
		return
	}

	err := h.Service.SoftDeleteRange(payload.Ids)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to delete the category",
		}, http.StatusInternalServerError)

		return
	}

	response.Response(w, response.View{
		Status:  true,
		Message: "Success",
	}, http.StatusOK)
}
