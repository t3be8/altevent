package user

import (
	"altevent/entity"
	"altevent/utils"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

type UserRepo struct {
	Db *gorm.DB
}

// Check users islogin with payload
func (ur *UserRepo) IsLogin(email, password string) (entity.User, error) {
	var u entity.User
	var pwd string

	query := "SELECT id, username, email, phone, password FROM users WHERE email = ?"

	if err := ur.Db.Raw(query, email).Scan(&u).Error; err != nil {
		log.Warn(err)
		return entity.User{}, errors.New("tidak dapat select data")
	}

	pwd = u.Password

	match, err := utils.CheckPasswordHash(password, pwd)
	if !match {
		log.Warn("Hash and password doesnt match")
		return u, err
	}

	log.Info()
	return u, nil
}

func (ur *UserRepo) Register(newUser entity.User) (entity.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		return entity.User{}, errors.New("tidak dapat insert data")
	}
	log.Info()
	return newUser, nil
}

func (ur *UserRepo) GetUserID(id uint) (entity.User, error) {
	arrUser := []entity.User{}

	if err := ur.Db.Where("id = ?", id).Find(&arrUser).Error; err != nil {
		log.Warn(err)
		return entity.User{}, errors.New("data select error")
	}

	if len(arrUser) == 0 {
		log.Warn("data tidak ditemukan")
		return entity.User{}, errors.New("data tidak ditemukan")
	}

	log.Info()
	return arrUser[0], nil
}

func (ur *UserRepo) UpdateUser(id uint, update entity.User) (entity.User, error) {
	var user entity.User
	if err := ur.Db.Where("id = ?", id).Updates(&update).Find(&user).Error; err != nil {
		log.Warn(err)
		return entity.User{}, errors.New("tidak bisa update data")
	}

	log.Info()
	return user, nil
}

func (ur *UserRepo) DeleteUser(id uint) (entity.User, error) {
	var user entity.User
	res, err := ur.GetUserID(id)
	if err != nil {
		return entity.User{}, err
	}

	if err := ur.Db.Delete(&user, "id = ?", id).Error; err != nil {
		log.Warn(err)
		return entity.User{}, errors.New("data select error")
	}
	log.Info()
	return res, nil
}

func (ur *UserRepo) GetMyEvent(id uint) ([]entity.Event, error) {
	var events []entity.Event
	result := ur.Db.Where("user_id = ?", id).Find(&events)
	if result.Error != nil {
		log.Warn(result.Error)
		return []entity.Event{}, errors.New("data select error")
	}

	if len(events) == 0 {
		log.Warn("data tidak ditemukan")
		return events, errors.New("data tidak ditemukan")
	}

	log.Info()
	return events, nil
}
