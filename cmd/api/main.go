package main

import (
	"github.com/AndresBC-Dev/tickets-api/v1/handler"
	"github.com/AndresBC-Dev/tickets-api/v1/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	eventRepository := repository.NewEventRepository(nil)

	r := gin.Default()

	handler.NewEventHandler(r, eventRepository)

	r.Run(":8030")
}
