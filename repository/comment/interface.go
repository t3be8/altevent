package comment

import "altevent/entity"

type IComment interface {
	CreateComment(comment entity.Comment) (entity.Comment, error)
	// GetSingleComment(evid uint) (entity.Comment, error)
	SelectAllComment(evid uint) ([]entity.Comment, error)
	UpdateComment(id, user_id uint, UpdateComment entity.Comment) (entity.Comment, error)
	DeleteComment(id, user_id uint) error
}
