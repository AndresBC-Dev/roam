package handler

import (
	"net/http"

	"github.com/AndresBC-Dev/tickets-api/v1/model"
	"github.com/AndresBC-Dev/tickets-api/v1/repository"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	repository repository.EventRepository
}

// Handler creator handler
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event model.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	createdEvent, err := h.repository.CreateOne(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, createdEvent)
}

func (h *EventHandler) GetAllEvents(c *gin.Context) {
	events, err := h.repository.GetMany(model.Event{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

// Constructor of eventHandler
func NewEventHandler(router *gin.Engine, repository *repository.EventRepository) {
	handler := EventHandler{
		repository: *repository,
	}

	r := router.Group("/event")
	{
		r.POST("/", handler.CreateEvent)
		r.GET("/", handler.GetAllEvents)
	}

}
