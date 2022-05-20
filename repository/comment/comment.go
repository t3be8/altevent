package comment

import (
	"altevent/entity"

	"gorm.io/gorm"
)

type CommentRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		Db: db,
	}
}

func (cr *CommentRepo) CreateComment(comment entity.Comment) (entity.Comment, error) {
	// var event entity.Event

	// if err := cr.Db.Where("id = ?", evid).First(&event).Error;err == nil {
	// 	return  event.ID, err
	// }

	if err := cr.Db.Create(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

// func (cr *CommentRepo) GetSingleComment(id uint) (entity.Comment, error) {
// 	var comment entity.Comment
// 	if err := cr.Db.Where("id = ?", id).First(&comment).Error; err != nil {
// 		return comment, err
// 	}
// 	return comment, nil
// }

func (cr *CommentRepo) SelectAllComment(evid uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := cr.Db.Where("event_id = ?", evid).Find(&comments).Limit(10).Error; err != nil {
		return comments, err
	}
	return comments, nil
}
func (cr *CommentRepo) UpdateComment(id, user_id uint, UpdateComment entity.Comment) (entity.Comment, error) {
	var comment entity.Comment

	if err := cr.Db.Where("id =? AND user_id =?", id, user_id).First(&comment).Updates(&UpdateComment).Find(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}
func (cr *CommentRepo) DeleteComment(id, user_id uint) error {
	var comment entity.Comment

	if err := cr.Db.Where("id = ? AND user_id = ?", id, user_id).First(&comment).Error; err != nil {
		return err
	}
	return nil
}
