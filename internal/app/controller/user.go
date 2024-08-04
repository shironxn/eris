package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/app/view"
	"github.com/shironxn/eris/internal/infrastructure/service"
)

type UserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) Login(ctx *gin.Context) {
	var req model.Login

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Login(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}

func (u *userController) Register(ctx *gin.Context) {
	var req model.Register

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Register(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusCreated, nil)
}

func (u *userController) GetAll(ctx *gin.Context) {
	data, err := u.service.GetAll()
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
	}

	var users []model.UserResponse
	for _, user := range data {
		users = append(users, model.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	view.JSON(ctx, http.StatusOK, users)
}

func (u *userController) GetByID(ctx *gin.Context) {
	var req model.UserUri

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	data, err := u.service.GetByID(req.ID)
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, model.UserResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}

func (u *userController) Update(ctx *gin.Context) {
	var req model.UserUpdate

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Update(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}

func (u *userController) Delete(ctx *gin.Context) {
	var req model.UserUri

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Delete(req.ID); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}
