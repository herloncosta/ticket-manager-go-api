package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/herloncosta/ticket-manager/internal/app/model"
	"github.com/herloncosta/ticket-manager/internal/app/service"
	"net/http"
	"strconv"
)

type TicketHandler struct {
	service service.TicketService
}

func NewTicketService(s service.TicketService) *TicketHandler {
	return &TicketHandler{service: s}
}

func (h *TicketHandler) Create(c *gin.Context) {
	var ticket model.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(&ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ticket.ID = id
	c.JSON(http.StatusCreated, ticket)
}

func (h *TicketHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	ticket, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No ticket found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) List(c *gin.Context) {
	tickets, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var ticket model.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.ID = id

	if err := h.service.Update(&ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
