package events

import "altevent/entity"

type IEvents interface {
	InsertEvent(newEvent entity.Event) (entity.Event, error)
	SelectEvent() ([]entity.Event, error)
	GetEventID(id uint) (entity.Event, error)
	UpdateEvent(id uint, update entity.Event) (entity.Event, error)
	DeleteEvent(id uint) (entity.Event, error)
}
