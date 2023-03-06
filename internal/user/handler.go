package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/jumaniyozov/grest/internal/apperror"
	"github.com/jumaniyozov/grest/internal/handlers"
	"github.com/jumaniyozov/grest/pkg/logging"
	"net/http"
)

const (
	usersURL      = "/users"
	singleUserURL = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, singleUserURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, singleUserURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, singleUserURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, singleUserURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.ErrNotFound
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("error creating user")
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "test uuid", "test", "US-000001")
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Update User By UUID"))
	return nil
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Partially User By UUID"))
	return nil
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Delete User By UUID"))
	return nil
}
