package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Index struct {
}

func (i Index) Index(c *gin.Context) {

	log.Println(c.Request)

	c.HTML(http.StatusOK, "index.html", nil)
}

func NewIndex() *Index {
	return &Index{}
}
