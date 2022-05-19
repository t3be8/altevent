package user

import (
	"altevent/delivery/middlewares"
	view "altevent/delivery/views"
	"altevent/delivery/views/req"
	"altevent/delivery/views/res"
	"altevent/entity"
	userRepo "altevent/repository/user"
	"altevent/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  userRepo.IUser
	Valid *validator.Validate
}

func New(repo userRepo.IUser, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUser req.RegisterRequest
		var resp res.UserResponse

		if err := c.Bind(&tmpUser); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, view.StatusInvalidRequest())
		}

		if err := uc.Valid.Struct(tmpUser); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.StatusValidate())
		}

		pwd := tmpUser.Password
		hash, _ := utils.HashPassword(pwd)

		newUser := entity.User{
			FullName: tmpUser.Fullname,
			Username: tmpUser.Username,
			Email:    tmpUser.Email,
			Phone:    tmpUser.Phone,
			Password: hash,
		}

		data, err := uc.Repo.Register(newUser)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		resp = res.UserResponse{
			ID:       data.ID,
			Username: data.Username,
			Email:    data.Email,
			Phone:    data.Phone,
		}

		log.Info("berhasil register")
		return c.JSON(http.StatusCreated, res.RegisterSuccess(resp))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var resp res.UserResponse
		var token string
		param := req.LoginRequest{}

		if err := c.Bind(&param); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.StatusBindData())
		}

		if err := uc.Valid.Struct(param); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.StatusValidate())
		}

		match, err := uc.Repo.IsLogin(param.Email, param.Password)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, view.StatusUnauthorized(err))
		}

		resp = res.UserResponse{
			ID:       match.ID,
			Username: match.Username,
			Email:    match.Email,
			Phone:    match.Phone,
		}

		res := res.LoginResponse{Data: resp, Token: token}

		if res.Token == "" {
			token, _ = middlewares.CreateToken(float64(match.ID), match.Username, match.Email)
			res.Token = token
			return c.JSON(http.StatusOK, view.StatusOK("Berhasil login!", res))
		}

		// c.SetCookie(&http.Cookie{
		// 	Name:    "token",
		// 	Value:   res.Token,
		// 	Expires: time.Now().Add(time.Hour * 1),
		// })

		return c.JSON(http.StatusOK, view.StatusOK("Berhasil login!", res))
	}
}

func (uc *UserController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, view.StatusIdConversion())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("Data not found"))
		}

		user, err := uc.Repo.GetUserID(uint(convID))

		if err != nil {
			log.Warn()
			return c.JSON(http.StatusNotFound, view.StatusNotFound("Data not found"))
		}
		response := res.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		}
		return c.JSON(http.StatusOK, view.StatusGetDatIdOK(response))
	}

}

func (uc *UserController) ShowMyEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, view.StatusIdConversion())
		}

		UserID := middlewares.ExtractTokenUserId(c)
		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("Data not found"))
		}

		result, err := uc.Repo.GetMyEvent(uint(UserID))
		if err != nil {
			log.Warn()
			return c.JSON(http.StatusNotFound, view.StatusNotFound("Data not found"))
		}

		var arrEvent []res.EventFullResponse
		for _, v := range result {
			event := res.EventFullResponse{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Rules:       v.Rules,
				Banner:      v.Banner,
				DueDate:     v.DueDate,
				BeginAt:     v.BeginAt,
				Location:    v.Location,
				Organizer:   v.Organizer,
				Ticket:      v.Ticket,
				Links:       v.Links,
				UserID:      v.UserID,
			}
			arrEvent = append(arrEvent, event)
		}

		return c.JSON(http.StatusOK, view.StatusGetDatIdOK(arrEvent))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUpdate req.UpdateUserReq

		if err := c.Bind(&tmpUpdate); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.StatusBindData())
		}

		if err := uc.Valid.Struct(tmpUpdate); err != nil {
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

		var pwd string
		pwd, err = utils.HashPassword(tmpUpdate.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		UpdateUser := entity.User{
			FullName: tmpUpdate.Fullname,
			Email:    tmpUpdate.Email,
			Phone:    tmpUpdate.Phone,
			Password: pwd,
		}

		user, err := uc.Repo.UpdateUser(uint(id), UpdateUser)

		if err != nil {
			log.Warn(err)
			notFound := "data tidak ditemukan"
			if err.Error() == notFound {
				return c.JSON(http.StatusNotFound, view.StatusNotFound(notFound))
			}
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}
		response := res.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		}
		return c.JSON(http.StatusOK, view.StatusUpdate(response))
	}

}

func (uc *UserController) Delete() echo.HandlerFunc {
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

		found, err := uc.Repo.GetUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("data tidak ditemukan"))
		}

		_, error := uc.Repo.DeleteUser(found.ID)

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
