package service

import (
	"github.com/herloncosta/ticket-manager/internal/app/model"
	"github.com/herloncosta/ticket-manager/internal/app/repository"
)

type TicketService interface {
	Create(ticket *model.Ticket) (int64, error)
	Get(id int64) (*model.Ticket, error)
	List() ([]model.Ticket, error)
	Update(ticket *model.Ticket) error
	Delete(id int64) error
}

type ticketService struct {
	repository repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) TicketService {
	return &ticketService{repository: repo}
}

func (s *ticketService) Create(ticket *model.Ticket) (int64, error) {
	ticket.Status = "OPEN"
	return s.repository.Create(ticket)
}

func (s *ticketService) Get(id int64) (*model.Ticket, error) {
	return s.repository.GetByID(id)
}

func (s *ticketService) List() ([]model.Ticket, error) {
	return s.repository.List()
}

func (s *ticketService) Update(ticket *model.Ticket) error {
	return s.repository.Update(ticket)
}

func (s *ticketService) Delete(id int64) error {
	return s.repository.Delete(id)
}
