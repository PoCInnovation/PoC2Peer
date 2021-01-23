package httpHost

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Host struct {
	*gin.Engine
	id string
	Port int
}

func NewHttpProvider(port int, id string) (*Host, error) {
	h := &Host{Engine: gin.Default(), id: id, Port: port}
	if err := h.applyRoutes(); err != nil {
		return nil, err
	}
	return h, nil
}

func (h *Host) applyRoutes() error {
	h.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	h.GET("/ID", func(c *gin.Context) {
		c.String(http.StatusOK, h.id)
	})
	return nil
}

func (h *Host) String() string {
	return h.id
}
