package repository

import (
	"database/sql"
	"github.com/herloncosta/ticket-manager/internal/app/model"
	"time"
)

type TicketRepository interface {
	Create(ticket *model.Ticket) (int64, error)
	GetByID(id int64) (*model.Ticket, error)
	List() ([]model.Ticket, error)
	Update(ticket *model.Ticket) error
	Delete(id int64) error
}

type ticketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) Create(ticket *model.Ticket) (int64, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO tickets (title, description, client_id, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	now := time.Now()

	res, err := stmt.Exec(ticket.Title, ticket.Description, ticket.ClientID, ticket.Status, now, now)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

func (r *ticketRepository) GetByID(id int64) (*model.Ticket, error) {
	row := r.db.QueryRow(`SELECT id, title, description, client_id, status, created_at, updated_at FROM tickets WHERE id = ?`, id)

	t := &model.Ticket{}
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.ClientID, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *ticketRepository) List() ([]model.Ticket, error) {
	rows, err := r.db.Query(`SELECT id, title, description, client_id, status, created_at, updated_at FROM tickets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []model.Ticket
	for rows.Next() {
		var t model.Ticket
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.ClientID, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func (r *ticketRepository) Update(ticket *model.Ticket) error {
	_, err := r.db.Exec(`
		UPDATE tickets SET title = ?, description = ?, client_id = ?, status = ?, updated_at = ? WHERE id = ?
	`, ticket.Title, ticket.Description, ticket.ClientID, ticket.Status, time.Now(), ticket.ID)
	return err
}

func (r *ticketRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM tickets WHERE id = ?`, id)
	return err
}
