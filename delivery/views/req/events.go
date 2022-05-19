package req

type CreateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rules       string `json:"rules"`
	Organizer   string `json:"organizer"`
	DueDate     string `json:"due_date"`
	BeginAt     string `json:"begin_at"`
	Location    string `json:"location"`
	Ticket      int    `json:"ticket"`
	Links       string `json:"links"`
	Banner      string `json:"banner"`
}

type UpdateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rules       string `json:"rules"`
	Organizer   string `json:"organizer"`
	DueDate     string `json:"due_date"`
	BeginAt     string `json:"begin_at"`
	Location    string `json:"location"`
	Ticket      int    `json:"ticket"`
	Links       string `json:"links"`
	Banner      string `json:"banner"`
}
