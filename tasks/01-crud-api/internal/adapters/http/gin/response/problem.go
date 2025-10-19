package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Problem struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status"`
}

func write(c *gin.Context, code int, typ, title, detail string) {
	c.JSON(code, Problem{Type: typ, Title: title, Detail: detail, Status: code})
}

func BadRequest(c *gin.Context, typ, detail string) {
	write(c, http.StatusBadRequest, typ, "Bad Request", detail)
}
func NotFound(c *gin.Context, typ, detail string) {
	write(c, http.StatusNotFound, typ, "Not Found", detail)
}
