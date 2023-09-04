package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Index struct{}

func (i Index) Index(c *gin.Context) {
	if strings.Contains(c.Request.Host, "danajx.com") {
		c.HTML(http.StatusOK, "dana_index.html", nil)
	} else {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func NewIndex() *Index {
	return &Index{}
}
