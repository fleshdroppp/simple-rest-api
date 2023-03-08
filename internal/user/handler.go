package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-rest-api/internal/app_error"
	"simple-rest-api/internal/handlers"
	"simple-rest-api/pkg/logging"
)

const (
	usersURL = "/users"
	userURL  = "/user/:id"
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
	router.HandlerFunc(http.MethodGet, usersURL, app_error.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, app_error.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, app_error.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, app_error.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, app_error.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, app_error.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	return app_error.ErrNotFound
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {

	return fmt.Errorf("this is api error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get user by uuid!"))
	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("update user!"))
	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("partially update user!"))
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("delete user!"))

	return nil
}
