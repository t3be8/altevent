package event

import (
	"altevent/delivery/middlewares"
	view "altevent/delivery/views"
	"altevent/delivery/views/req"
	"altevent/delivery/views/res"
	"altevent/entity"
	eventRepo "altevent/repository/events"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type EventController struct {
	Repo  eventRepo.IEvents
	Valid *validator.Validate
}

func New(repo eventRepo.IEvents, valid *validator.Validate) *EventController {
	return &EventController{
		Repo:  repo,
		Valid: valid,
	}
}

func (ec *EventController) InsertEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpEvent req.CreateEventRequest
		var resp res.EventResponse

		if err := c.Bind(&tmpEvent); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, view.StatusInvalidRequest())
		}

		if err := ec.Valid.Struct(tmpEvent); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.StatusValidate())
		}

		newEvent := entity.Event{
			Title:       tmpEvent.Title,
			Description: tmpEvent.Description,
			Rules:       tmpEvent.Rules,
			Organizer:   tmpEvent.Organizer,
			DueDate:     tmpEvent.DueDate,
			BeginAt:     tmpEvent.BeginAt,
			Location:    tmpEvent.Location,
			Ticket:      tmpEvent.Ticket,
			Links:       tmpEvent.Links,
			Banner:      tmpEvent.Banner,
		}

		data, err := ec.Repo.Create(newEvent)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		resp = res.EventResponse{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			Rules:       data.Rules,
			Organizer:   data.Organizer,
			Location:    data.Location,
			Ticket:      data.Ticket,
			Links:       data.Links,
		}

		log.Info("berhasil register")
		return c.JSON(http.StatusCreated, res.CreateEventSuccess(resp))
	}
}

func (ec *EventController) GetEventById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, view.StatusIdConversion())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("data user tidak ditemukan"))
		}

		event, err := ec.Repo.GetEventID(uint(convID))

		if err != nil {
			log.Warn()
			return c.JSON(http.StatusNotFound, view.StatusNotFound("data user tidak ditemukan"))
		}
		response := res.EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Rules:       event.Rules,
			Organizer:   event.Organizer,
			Location:    event.Location,
			Ticket:      event.Ticket,
			Links:       event.Links,
		}
		return c.JSON(http.StatusOK, view.StatusGetDatIdOK(response))
	}

}

func (ec *EventController) SelectEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id := c.Param("id")

		// convID, err := strconv.Atoi(id)
		res, err := ec.Repo.GetEvent()

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil select Event")
		return c.JSON(http.StatusOK, res)
	}

}

func (ec *EventController) UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUpdate req.UpdateEventRequest

		if err := c.Bind(&tmpUpdate); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.StatusBindData())
		}

		if err := ec.Valid.Struct(tmpUpdate); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.StatusValidate())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.StatusIdConversion())
		}

		UserID := middlewares.ExtractTokenUserId(c)
		if UserID != float64(id) {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("token tidak ditemukan"))
		}

		updateEvent := entity.Event{
			Title:       tmpUpdate.Title,
			Description: tmpUpdate.Description,
			Rules:       tmpUpdate.Rules,
			Organizer:   tmpUpdate.Organizer,

			Location: tmpUpdate.Location,
			Ticket:   tmpUpdate.Ticket,
			Links:    tmpUpdate.Links,
			Banner:   tmpUpdate.Banner,
		}

		event, err := ec.Repo.UpdateEvent(uint(id), updateEvent)

		if err != nil {
			log.Warn(err)
			notFound := "data tidak ditemukan"
			if err.Error() == notFound {
				return c.JSON(http.StatusNotFound, view.StatusNotFound(notFound))
			}
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}
		response := res.EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Rules:       event.Rules,
			Organizer:   event.Organizer,
			Location:    event.Location,
			Ticket:      event.Ticket,
			Links:       event.Links,
		}
		return c.JSON(http.StatusOK, view.StatusUpdate(response))
	}

}

func (ec *EventController) DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.StatusIdConversion())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("data tidak ditemukan"))
		}

		found, err := ec.Repo.GetEventID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("data tidak ditemukan"))
		}

		_, error := ec.Repo.DeleteEvent(found.ID)

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
