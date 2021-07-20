package service

import (
	"log"

	"github.com/gin-gonic/gin"
)

type shortHandler struct {
	service Service
}

// func to intialize new handler
func NewShortHandler(service Service) *shortHandler {
	return &shortHandler{service}
}

func (h *shortHandler) CreateNewShortLink(c *gin.Context) {
	var input ShortLinkInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var userID = c.GetHeader("Authorization")

	shorted, err := h.service.Create(input.Link, userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, shorted)
}

func (h *shortHandler) RedirectLongLink(c *gin.Context) {
	shortUrl := c.Params.ByName("short_link")

	log.Println(shortUrl)

	shorted, err := h.service.FindLongLink(shortUrl)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(shorted.LongLink)

	c.Redirect(302, shorted.LongLink)
}

func (h *shortHandler) UpdateShortLink(c *gin.Context) {
	id := c.Params.ByName("short_link_id")

	var update UpdateShortLinkInput

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.Update(id, update.ShortLink)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updated)
}
