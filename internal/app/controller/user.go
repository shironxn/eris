package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/app/view"
	"github.com/shironxn/eris/internal/infrastructure/service"
)

type UserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) Login(c *gin.Context) {
	var req model.Login

	if err := c.ShouldBind(&req); err != nil {
		view.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Login(req); err != nil {
		view.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(c, http.StatusOK, nil)
}

func (u *userController) Register(c *gin.Context) {
	var req model.Register

	if err := c.ShouldBind(&req); err != nil {
		view.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Register(req); err != nil {
		view.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(c, http.StatusCreated, nil)
}

func (u *userController) GetAll(c *gin.Context) {
	data, err := u.service.GetAll()
	if err != nil {
		view.JSON(c, http.StatusInternalServerError, err.Error())
	}

	view.JSON(c, http.StatusOK, data)
}

func (u *userController) GetByID(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindUri(&req); err != nil {
		view.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := u.service.GetByID(req.ID)
	if err != nil {
		view.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(c, http.StatusOK, data)
}

func (u *userController) Update(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindUri(&req); err != nil {
		view.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		view.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Update(req); err != nil {
		view.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(c, http.StatusOK, nil)
}

func (u *userController) Delete(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindUri(&req); err != nil {
		view.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Delete(req.ID); err != nil {
		view.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(c, http.StatusOK, nil)
}
