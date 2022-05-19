package events

import (
	"altevent/entity"

	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *EventRepo {
	return &EventRepo{
		Db: db,
	}
}

type EventRepo struct {
	Db *gorm.DB
}

func (er *EventRepo) Create(newEvent entity.Event) (entity.Event, error) {
	if err := er.Db.Create(&newEvent).Error; err != nil {
		return entity.Event{}, errors.New("tidak dapat insert data")
	}
	log.Info()
	return newEvent, nil
}

func (er *EventRepo) GetEvent() ([]entity.Event, error) {
	arrEvent := []entity.Event{}

	if err := er.Db.Find(&arrEvent).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select")
	}

	if len(arrEvent) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrEvent, nil
}

func (er *EventRepo) GetEventID(id uint) (entity.Event, error) {
	arrEvent := []entity.Event{}

	if err := er.Db.Where("id = ?", id).Find(&arrEvent).Error; err != nil {
		log.Warn(err)
		return entity.Event{}, errors.New("tidak bisa select data")
	}

	if len(arrEvent) == 0 {
		log.Warn("data tidak ditemukan")
		return entity.Event{}, errors.New("data tidak ditemukan")
	}

	log.Info()
	return arrEvent[0], nil
}

func (er *EventRepo) UpdateEvent(id uint, update entity.Event) (entity.Event, error) {
	var event entity.Event
	if err := er.Db.Where("id = ?", id).Updates(&update).Find(&event).Error; err != nil {
		log.Warn(err)
		return entity.Event{}, errors.New("tidak bisa update data")
	}

	log.Info()
	return event, nil
}

func (er *EventRepo) DeleteEvent(id uint) (entity.Event, error) {
	var event []entity.Event
	res, err := er.GetEventID(id)
	if err != nil {
		return entity.Event{}, err
	}

	if err := er.Db.Delete(&event, "id = ?", id).Error; err != nil {
		log.Warn(err)
		return entity.Event{}, errors.New("tidak bisa delete data")
	}
	log.Info()
	return res, nil
}
