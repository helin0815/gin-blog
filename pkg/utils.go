package pkg

import (
	"net/http"
	"path"
	"time"

	"github.com/gin-blog/faye-wong/model"
	"github.com/gin-gonic/gin"
)

func GetCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02-15:04:05")
	return currentTime
}

func GetCurrentDatabaseTime() string {
	currentTime := time.Now().Format("2006-01-02-15:04:05")
	return currentTime
}

func JSONServerError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, NewResponse(http.StatusInternalServerError, nil, "", msg))
	c.Abort()
}

// JSONOk returns error with http.StatusOK
func JSONOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, NewResponse(0, map[string]interface{}{}, "", msg))
	c.Abort()
}

func NewResponse(code int, data interface{}, eCode, msg string) *model.RestResponse {
	return &model.RestResponse{
		Code: code,
		Data: data,
		Errors: []model.Error{
			{eCode, msg},
		},
	}
}

func GenerateErrMsg(msg, errMsg string) string {
	return path.Join(msg, errMsg)
}
