package response

import (
	"net/http"
	"os"
	"server/internal/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
}

func JSON(c *gin.Context, code int, data interface{}, err string) {
	c.JSON(code, Response{
		Data:       data,
		Error:      err,
		Status:     http.StatusText(code),
		StatusCode: code,
	})
}

func OK(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, data, "")
}

func Created(c *gin.Context, data interface{}) {
	JSON(c, http.StatusCreated, data, "")
}

func Internal(c *gin.Context, err string) {
	JSON(c, http.StatusInternalServerError, nil, err)
}

func HandleError(c *gin.Context, err error) {

	appErr, ok := errors.IsAppError(err)
	env := os.Getenv("APP_ENV")

	if !ok {
		if env == "PROD" {
			c.Redirect(http.StatusFound, "/error")
			return
		}
		Internal(c, "internal server error")
		return
	}

	if env == "PROD" && appErr.StatusCode >= 500 {
		c.Redirect(http.StatusFound, "/error")
		return
	}

	JSON(c, appErr.StatusCode, nil, appErr.Message)
}