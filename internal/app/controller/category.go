package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/app/view"
	"github.com/shironxn/eris/internal/infrastructure/service"
)

type CategoryController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{
		service: service,
	}
}

func (c *categoryController) Create(ctx *gin.Context) {
	var req model.CategoryCreate

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.Create(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusCreated, nil)
}

func (c *categoryController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var categories []model.CategoryResponse
	for _, category := range data {
		categories = append(categories, model.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	view.JSON(ctx, http.StatusOK, categories)
}

func (c *categoryController) GetByID(ctx *gin.Context) {
	var req model.CategoryUri

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	data, err := c.service.GetByID(req.ID)
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, model.CategoryResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}

func (c *categoryController) Update(ctx *gin.Context) {
	var req model.CategoryUpdate

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.Update(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}

func (c *categoryController) Delete(ctx *gin.Context) {
	var req model.CategoryUri

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.Delete(req.ID); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}
