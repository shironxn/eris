package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/app/view"
	"github.com/shironxn/eris/internal/infrastructure/service"
	"github.com/shironxn/eris/internal/infrastructure/util"
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
	jwt     util.JWT
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

	data, err := u.service.Login(req)
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := u.jwt.GenerateAccessToken(data.ID)
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := u.jwt.GenerateRefreshToken(data.ID)
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("access-token", accessToken, int(time.Now().Add(10*time.Minute).Unix()), "/", "localhost", false, true)
	ctx.SetCookie("refresh-token", refreshToken, int(time.Now().Add(24*time.Hour).Unix()), "/", "localhost", false, true)

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
