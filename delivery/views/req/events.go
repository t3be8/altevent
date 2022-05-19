package req

import "time"

type CreateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rules       string    `json:"rules"`
	Organizer   string    `json:"organizer"`
	DueDate     float64   `json:"duedate"`
	BeginAt     time.Time `json:"begin"`
	Location    string    `json:"location"`
	Ticket      int       `json:"ticket"`
	Links       string    `json:"links"`
	Banner      string    `json:"banner"`
}

type UpdateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rules       string    `json:"rules"`
	Organizer   string    `json:"organizer"`
	DueDate     float64   `json:"duedate"`
	BeginAt     time.Time `json:"begin"`
	Location    string    `json:"location"`
	Ticket      int       `json:"ticket"`
	Links       string    `json:"links"`
	Banner      string    `json:"banner"`
}
