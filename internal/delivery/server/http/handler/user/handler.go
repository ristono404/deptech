package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	config "github.com/ristono404/deptech/internal/config"
	event "github.com/ristono404/deptech/internal/event"
	auth "github.com/ristono404/deptech/internal/pkg/auth"
	pkgRequest "github.com/ristono404/deptech/internal/pkg/request"
	response "github.com/ristono404/deptech/internal/pkg/response"
	userUsecase "github.com/ristono404/deptech/internal/usecase/user"
)

type Handler struct {
	Service userUsecase.Service
	Config  config.Config
}

func New(service userUsecase.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

/*
IsISO8601Date function to check parameter pattern for valid ISO8601 Date
*/
func IsISO8601Date(fl validator.FieldLevel) bool {
	format := fmt.Sprintf("%s 23:59:59", fl.Field().String())
	_, err := time.Parse("2006-01-02 15:04:05", format)
	return err == nil
}

/*CustomValidator struct is for storing the custom validator that will be registered to echo server */
type CustomValidator struct {
	Validator *validator.Validate
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	page, limit, offset, from, to, err := pkgRequest.List(r.URL.Query(), h.Config.PerPage)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to get user",
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
			Message: "failed to get user",
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

	var body event.UserCreated
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
	validator := validator.New()
	validator.RegisterValidation("isodate", IsISO8601Date)

	if err := validator.Struct(body); err != nil {
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

	var body event.UserCreated
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

	validator := validator.New()
	validator.RegisterValidation("isodate", IsISO8601Date)

	if err := validator.Struct(body); err != nil {
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
			Message: "failed to get user",
		}, http.StatusInternalServerError)
		return
	}

	data, err := h.Service.Read(id)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to read the user",
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body event.UserLogin
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

	data, err := h.Service.Login(body)
	if err != nil {
		log.Println(err)

		response.Response(w, response.View{
			Message: "failed to  login",
		}, http.StatusInternalServerError)

		return
	}

	if data == nil {
		response.Response(w, response.View{
			Message:      "Error",
			ErrorMessage: "email/password wrong",
		}, http.StatusBadRequest)
		return
	}

	token, _ := auth.GenerateJWT(data.Email)

	response.Response(w, response.View{
		Status: true,
		Data:   token,
	}, http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	auth.Logout(token)
	response.Response(w, response.View{
		Status: true,
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
			Message: "failed to delete the user",
		}, http.StatusInternalServerError)

		return
	}

	response.Response(w, response.View{
		Status:  true,
		Message: "Success",
	}, http.StatusOK)
}
