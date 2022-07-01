package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
}

func (i Index) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func NewIndex() *Index {
	return &Index{}
}
