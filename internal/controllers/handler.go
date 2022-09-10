package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/bsgilber/heart-beat/internal/models"
	"github.com/gin-gonic/gin"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	endpointRepo models.EndpointRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(endpointRepo models.EndpointRepository) *BaseHandler {
	return &BaseHandler{
		endpointRepo: endpointRepo,
	}
}

func (h *BaseHandler) Ping(c *gin.Context, name string) {
	endpoint, err := h.endpointRepo.FindByName(name)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("the key [%s] does not exist.\n", name))
		return
	}

	req, err := http.NewRequest(http.MethodGet, endpoint.URL, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	defer res.Body.Close()

	c.String(http.StatusOK, "[%s] had status code [%d]\n", name, res.StatusCode)
}

func (h *BaseHandler) Register(c *gin.Context) {
	var endpoint models.Endpoint

	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ep, _ := h.endpointRepo.FindByName(endpoint.Name)
	if ep != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "name must be unique"})
		return
	}

	u, err := url.Parse(endpoint.URL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		c.String(http.StatusBadRequest, "invalid url, url must contain scheme, host, and/or path")
		return
	}

	if err := h.endpointRepo.Save(endpoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("[%s] has been added to the cache.", endpoint.Name))
}

func (h *BaseHandler) SinglePingPrep(c *gin.Context) {
	name := c.Param("name")

	h.Ping(c, name)
}

func (h *BaseHandler) AllPingPrep(c *gin.Context) {
	for endpoint := range h.Endpoints {
		h.Ping(c, endpoint.Name)
	}
}

func (h *BaseHandler) ListRegistered(c *gin.Context) {
	b, err := json.Marshal(cache)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, string(b))
}
