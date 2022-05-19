package res

import (
	"net/http"
)

type EventResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rules       string `json:"rules"`
	Organizer   string `json:"organizer"`
	Location    string `json:"location"`
	Ticket      int    `json:"ticket"`
	Links       string `json:"links"`
}

type EventFullResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rules       string `json:"rules"`
	Banner      string `json:"banner"`
	DueDate     string `json:"due_date"`
	BeginAt     string `json:"begin_at"`
	Organizer   string `json:"organizer"`
	Location    string `json:"location"`
	Ticket      int    `json:"ticket"`
	Links       string `json:"links"`
	UserID      uint   `json:"user_id"`
}

func CreateEventSuccess(data EventResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil membuat event baru",
		"status":  true,
		"data":    data,
	}
}

func SelectEventSuccess(data EventResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil melihat event",
		"status":  true,
		"data":    data,
	}
}
