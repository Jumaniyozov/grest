package user

import (
	"github.com/julienschmidt/httprouter"
	"github.com/jumaniyozov/grest/internal/handlers"
	"net/http"
)

const (
	usersURL      = "/users"
	singleUserURL = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(singleUserURL, h.GetUserByUUID)
	router.PUT(singleUserURL, h.UpdateUser)
	router.PATCH(singleUserURL, h.PartiallyUpdateUser)
	router.DELETE(singleUserURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Get User List"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create User"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Get User By UUID"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Update User By UUID"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Partially User By UUID"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Delete User By UUID"))
}
