package res

type CommentResponse struct {
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Comment string `json:"comment"`
}
