package user

import "altevent/entity"

type IUser interface {
	Register(newUser entity.User) (entity.User, error)
	IsLogin(email, password string) (entity.User, error)
	GetUserID(id uint) (entity.User, error)
	UpdateUser(id uint, update entity.User) (entity.User, error)
	DeleteUser(id uint) (entity.User, error)
	GetMyEvent(id uint) ([]entity.Event, error)
}
