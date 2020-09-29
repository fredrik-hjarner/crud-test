package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fredrik-hjarner/crud-test/api"
	"github.com/fredrik-hjarner/crud-test/requestmodels"
	"github.com/fredrik-hjarner/crud-test/storage"
	"github.com/gorilla/mux"
)

// UserHandler ...
type UserHandler struct {
	*api.BaseHandler

	/* app *switchr.Application */
}

// NewUserHandler ...
func NewUserHandler() *UserHandler {
	return &UserHandler{
		BaseHandler: api.NewBaseHandler( /* app.Logger.WithField("handler", "cmsRpcHandler") */ ),
		/* app:         app, */
	}
}

// GetAll ...
func (handler *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	/* log := handler.NewLogEntry(r.Context()).WithField("operation", "DeletePage") */

	api.JsonResponse(w, http.StatusOK, map[string]interface{}{
		"data": storage.Users,
		"meta": api.CollectionMeta{
			Total:  len(storage.Users),
			Offset: 1, // TODO: criteria.Offset,
			Limit:  1, // TODO: criteria.Limit,
		},
	})
}

// Create ...
func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := &requestmodels.CreateUserRequest{}

	if err := handler.ParseRequestBody(r.Body, data); err != nil {
		api.WriteBadRequestResponse(w, "Failed to parse request body")
		return
	}

	ctx := context.Background() // TODO: No idea what a context is.

	if valid, _ := handler.Validate(ctx, w, data); !valid {
		return
	}

	user := data.ToUser()
	storage.AddUser(user)
	api.JsonResponse(w, http.StatusCreated, user)
}

// Update ...
func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	data := &requestmodels.UpdateUserRequest{}

	if err := handler.ParseRequestBody(r.Body, data); err != nil {
		api.WriteBadRequestResponse(w, "Failed to parse request body")
		return
	}

	ctx := context.Background() // TODO: No idea what a context is.

	if valid, _ := handler.Validate(ctx, w, data); !valid {
		return
	}

	// Check if user exists, else not found.
	vars := mux.Vars(r)
	id := vars["id"]

	_, getUserByIDError := storage.GetUserByID(id)

	if getUserByIDError != nil {
		api.WriteNotFoundResponse(w, fmt.Sprintf("User with id '%v' does not exist", id))
		return
	}

	newUser := data.ToUser(id)
	storage.ReplaceUser(id, newUser)
	api.JsonResponse(w, http.StatusCreated, newUser)
}
