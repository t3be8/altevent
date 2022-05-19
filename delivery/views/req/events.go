package req

type CreateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rules       string `json:"rules"`
	Organizer   string `json:"organizer"`
	DueDate     string `json:"duedate"`
	BeginAt     string `json:"begin"`
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
	DueDate     string `json:"duedate"`
	BeginAt     string `json:"begin"`
	Location    string `json:"location"`
	Ticket      int    `json:"ticket"`
	Links       string `json:"links"`
	Banner      string `json:"banner"`
}
