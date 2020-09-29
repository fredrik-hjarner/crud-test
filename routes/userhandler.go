package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/crud-test/api"
	"github.com/fredrik-hjarner/crud-test/requestmodels"
	"github.com/fredrik-hjarner/crud-test/storage"
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
	log.Printf("data: %v", data)

	ctx := context.Background() // TODO: No idea what a context is.
	// ctx = context.WithValue(ctx, "CurrentUser", nil)

	if valid, _ := handler.Validate(ctx, w, data); !valid {
		return
	}

	user := data.ToUser()
	storage.AddUser(user)
	api.JsonResponse(w, http.StatusCreated, user)
}
