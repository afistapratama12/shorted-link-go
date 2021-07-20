package service

import (
	"log"

	"github.com/gin-gonic/gin"
)

type shortHandler struct {
	service Service
}

func NewShortHandler(service Service) *shortHandler {
	return &shortHandler{service}
}

func (h *shortHandler) CreateNewShortLink(c *gin.Context) {
	var input ShortLinkInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	shorted, err := h.service.Create(input.Link)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, shorted)
}

func (h *shortHandler) ShowAllShortLink(c *gin.Context) {
	shorteds, _ := h.service.GetAll()

	c.JSON(200, shorteds)
}

func (h *shortHandler) RedirectLongLink(c *gin.Context) {
	param := c.Params.ByName("short_link")

	log.Println(param)

	shorted, err := h.service.FindLongLink(param)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(shorted.LongLink)

	// c.Redirect(http.StatusMovedPermanently, shorted.LongLink)

	// location := url.URL{}
	// c.Request.URL
}
