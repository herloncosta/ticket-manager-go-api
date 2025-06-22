package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/herloncosta/ticket-manager/internal/app/handler"
	"github.com/herloncosta/ticket-manager/internal/app/repository"
	"github.com/herloncosta/ticket-manager/internal/app/service"
	"net/http"
)

func Setup(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketService(ticketService)

	t := r.Group("/tickets")
	{
		t.POST("/", ticketHandler.Create)
		t.GET("/", ticketHandler.List)
		t.GET("/:id", ticketHandler.Get)
		t.PUT("/:id", ticketHandler.Update)
		t.DELETE("/:id", ticketHandler.Delete)
	}

	return r
}
