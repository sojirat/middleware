package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode" validate:"required"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
	XCSRFToken string      `json:"csrf,omitempty"`
}

// HTTP status code 200
func StatusOK(c *gin.Context, data interface{}) {
	SendResponse(c, http.StatusOK, "", data)
}

// HTTP status code 201
func StatusCreated(c *gin.Context, data interface{}) {
	SendResponse(c, http.StatusCreated, "", data)
}

// HTTP status code 202
func StatusAccepted(c *gin.Context, data interface{}) {
	SendResponse(c, http.StatusAccepted, "", data)
}

// HTTP status code 204
func StatusNoContent(c *gin.Context, message string) {
	SendResponse(c, http.StatusNoContent, message, nil)
}

// HTTP status code 307
func StatusTemporaryRedirect(c *gin.Context, message string, data interface{}) {
	SendResponse(c, http.StatusTemporaryRedirect, message, data)
}

// HTTP status code 400
func StatusBadRequest(c *gin.Context, err error) {
	SendResponse(c, http.StatusBadRequest, err.Error(), nil)
}

// HTTP status code 401
func StatusUnauthorized(c *gin.Context, message string) {
	SendResponse(c, http.StatusUnauthorized, message, nil)
}

// HTTP status code 403
func StatusForbidden(c *gin.Context) {
	SendResponse(c, http.StatusForbidden, "forbidden access", nil)
}

// HTTP status code 404
func StatusNotFound(c *gin.Context, message string) {
	SendResponse(c, http.StatusNotFound, message, nil)
}

// HTTP status code 417
func StatusExpectationFailed(c *gin.Context, message string) {
	SendResponse(c, http.StatusExpectationFailed, message, nil)
}

// HTTP status code 423
func StatusLocked(c *gin.Context, message string) {
	SendResponse(c, http.StatusLocked, message, nil)
}

// HTTP status code 428
func StatusPreconditionFailed(c *gin.Context, message string) {
	SendResponse(c, http.StatusPreconditionFailed, message, nil)
}

// HTTP status code 500
func StatusInternalServerError(c *gin.Context, err error) {
	SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
}

// HTTP status code 503
func StatusServiceUnavailable(c *gin.Context, service string, err error) {
	msg := fmt.Sprintf("%s unavailable: %s", service, err.Error())
	SendResponse(c, http.StatusServiceUnavailable, msg, nil)
}

// Send Response
func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		StatusCode: status,
		Message:    strings.ToLower(message),
		Data:       data,
		XCSRFToken: c.GetHeader("X-CSRF-Token"),
	})
}
