package comment

import (
	"altevent/delivery/middlewares"
	view "altevent/delivery/views"
	"altevent/delivery/views/req"
	"altevent/delivery/views/res"
	"altevent/entity"
	comRepo "altevent/repository/comment"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CommentController struct {
	Repo  comRepo.IComment
	Valid *validator.Validate
}

func New(repo comRepo.IComment, valid *validator.Validate) *CommentController {
	return &CommentController{
		Repo:  repo,
		Valid: valid,
	}
}

func (cc *CommentController) PostComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var tmpComment req.ReqComment
		if err := c.Bind(&tmpComment); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.StatusBindData())
		}

		if err := cc.Valid.Struct(&tmpComment); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.StatusValidate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		insert := entity.Comment{
			UserID:  uint(UserID),
			EventID: uint(id),
			Comment: tmpComment.Comment,
		}
		result, err := cc.Repo.CreateComment(insert)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, view.StatusCreated("Success Post Comment!", result))
	}
}

func (cc *CommentController) SelectAllInEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		result, err := cc.Repo.SelectAllComment(uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, view.StatusNotFound("Data not found"))
		}

		var arrCom []res.CommentResponse
		for _, v := range result {
			comment := res.CommentResponse{
				UserID:  v.UserID,
				EventID: v.EventID,
				Comment: v.Comment,
			}
			arrCom = append(arrCom, comment)
		}
		return c.JSON(http.StatusOK, view.StatusOK("Success Get Data", arrCom))
	}
}

func (cc *CommentController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUpdate req.ReqComment
		cid, _ := strconv.Atoi(c.Param("id"))
		user_id := middlewares.ExtractTokenUserId(c)

		if err := c.Bind(&tmpUpdate); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.StatusBindData())
		}

		update := entity.Comment{
			Comment: tmpUpdate.Comment,
		}

		result, err := cc.Repo.UpdateComment(uint(cid), uint(user_id), update)
		if err != nil {
			c.JSON(http.StatusNotFound, view.StatusNotFound("Data not found"))
		}

		return c.JSON(http.StatusOK, view.StatusUpdate(result))
	}
}

func (cc *CommentController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := middlewares.ExtractTokenUserId(c)
		err := cc.Repo.DeleteComment(uint(id), uint(user_id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
