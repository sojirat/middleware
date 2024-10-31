package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Response struct {
	Status     int         `json:"statusCode" validate:"required"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
	XCSRFToken string      `json:"csrf,omitempty"`
}

var validate = validator.New()

func WithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 10*time.Second)
}

func ValidateAndBind(c *gin.Context, v interface{}) bool {
	if err := c.BindJSON(v); err != nil {
		HandleBadRequest(c, err)
		return false
	}

	if validationErr := validate.Struct(v); validationErr != nil {
		HandleValidationError(c, validationErr)
		return false
	}
	return true
}

func HandleUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{Status: 0, Message: strings.ToLower(fmt.Sprintf("%v", message)), XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleValidationError(c *gin.Context, validationErr error) {
	c.JSON(http.StatusBadRequest, Response{Status: 0, Message: validationErr.Error(), XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleInternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{Status: 0, Message: err.Error(), XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Response{Status: 0, Message: err.Error(), XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{Status: 0, Message: strings.ToLower(fmt.Sprintf("%v", message)), XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{Status: 1, Data: data, XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Status: 1, Data: data, XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleNoContent(c *gin.Context, message string) {
	c.JSON(http.StatusNoContent, Response{Status: 1, Data: message, XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{Status: 0, Message: "unauthorized", XCSRFToken: c.GetHeader("X-CSRF-Token")})
}

func HandleServiceUnavailable(c *gin.Context, err error) {
	msg := "ServiceUnavailable: " + err.Error()
	c.JSON(http.StatusServiceUnavailable, Response{Status: 0, Message: msg, XCSRFToken: c.GetHeader("X-CSRF-Token")})
}
