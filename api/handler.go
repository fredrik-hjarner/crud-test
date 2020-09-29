package api

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"context"

	"net/http"

	"github.com/interactive-solutions/govalidator"
	"github.com/pkg/errors"
)

type BaseHandler struct {
	/* logger logrus.FieldLogger */
}

// NewBaseHandler ...
func NewBaseHandler( /* logger logrus.FieldLogger */ ) *BaseHandler {
	return &BaseHandler{
		/* logger: logger, */
	}
}

// ParseRequestBody ...
func (handler *BaseHandler) ParseRequestBody(body io.Reader, obj interface{}) error {
	bodyContent, err := ioutil.ReadAll(body)
	if err != nil {
		/* handler.logger.
		WithError(err).
		Error("Failed to read the request body, multiple calls to ParseRequestBody ?") */

		return errors.Wrapf(err, "Failed to read body")
	}

	if err = json.Unmarshal(bodyContent, obj); err != nil {
		/* handler.logger.
		WithField("body", string(bodyContent)).
		WithError(err).
		Error("Failed to decode json request body") */

		return errors.Wrapf(err, "Failed to parse body")
	}

	return nil
}

// NewLogEntry ...
/* func (handler *BaseHandler) NewLogEntry(ctx context.Context) logrus.FieldLogger {
	log := handler.logger

	if requestId, ok := ctx.Value("requestId").(int); ok {
		log = log.WithField("ctxRequestId", requestId)
	}

	return log
} */

// Validate ...
func (handler *BaseHandler) Validate(ctx context.Context, w http.ResponseWriter, data interface{}) (bool, error) {
	if valid, errorMap, err := govalidator.ValidateStruct(ctx, data); !valid {
		if err != nil {
			WriteInternalServerErrorResponse(w, "Unknown error validating data")
			return false, err
		}

		WriteValidationErrorResponse(w, errorMap, nil)
		return false, nil
	}

	return true, nil
}
