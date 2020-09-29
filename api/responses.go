package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ErrorResponseBody struct {
	// Note: StatusCode is the only snake cased json response we have in our entire api
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type CollectionMeta struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

var (
	ForbiddenResponseBody = ErrorResponseBody{
		StatusCode: http.StatusForbidden,
		Message:    "You do not have access to this endpoint, this request has been logged",
	}
)

type ValidationErrorMap map[string]interface{}
type ValidationMetaMap map[string]interface{}

func WriteValidationErrorResponse(w http.ResponseWriter, errorMap ValidationErrorMap, metaMap ValidationMetaMap) {
	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		StatusCode uint               `json:"status_code"`
		Message    string             `json:"message"`
		Errors     ValidationErrorMap `json:"errors"`
		Meta       ValidationMetaMap  `json:"meta"`
	}{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "validation errors",
		Errors:     errorMap,
		Meta:       metaMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(body)
}

func WriteForbiddenError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(&ForbiddenResponseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create response: %s", err)))
		return
	}

	w.WriteHeader(http.StatusForbidden)
	w.Write(body)
}

func JsonResponse(w http.ResponseWriter, status int, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode json response"))

		return errors.Wrap(err, "Failed to encode json payload")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)

	return nil
}

func WriteNotFoundResponse(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Detail     string `json:"detail"`
		HttpStatus int    `json:"httpStatus"`
		Title      string `json:"title"`
	}{
		Detail:     detail,
		HttpStatus: http.StatusNotFound,
		Title:      "Not found",
	}

	body, err := json.Marshal(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create response: %s", err)))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(body)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Detail     string `json:"detail"`
		HttpStatus int    `json:"httpStatus"`
		Title      string `json:"title"`
	}{
		Detail:     detail,
		HttpStatus: http.StatusInternalServerError,
		Title:      "Internal server error",
	}

	body, err := json.Marshal(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create response: %s", err)))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(body)
}

func WriteBadRequestResponse(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Detail     string `json:"detail"`
		HttpStatus int    `json:"httpStatus"`
		Title      string `json:"title"`
	}{
		Detail:     detail,
		HttpStatus: http.StatusBadRequest,
		Title:      "Bad request",
	}

	body, err := json.Marshal(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create response: %s", err)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(body)
}

func WriteUnauthorizedResponse(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Detail     string `json:"detail"`
		HttpStatus int    `json:"httpStatus"`
		Title      string `json:"title"`
	}{
		Detail:     detail,
		HttpStatus: http.StatusUnauthorized,
		Title:      "Bad request",
	}

	body, err := json.Marshal(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create response: %s", err)))
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write(body)
}
