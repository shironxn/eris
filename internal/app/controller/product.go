package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/app/view"
	"github.com/shironxn/eris/internal/infrastructure/service"
)

type ProductController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type productController struct {
	service service.ProductService
}

func NewProducController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (p *productController) Create(ctx *gin.Context) {
	var req model.ProductCreate

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.service.Create(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusCreated, nil)
}

func (p *productController) GetAll(ctx *gin.Context) {
	data, err := p.service.GetAll()
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var products []model.ProductResponse
	for _, product := range data {
		products = append(products, model.ProductResponse{
			ID:          product.ID,
			Name:        product.Description,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			UserID:      product.UserID,
			CategoryID:  product.CategoryID,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	view.JSON(ctx, http.StatusOK, products)
}

func (p *productController) GetByID(ctx *gin.Context) {
	var req model.ProductUri

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	data, err := p.service.GetByID(req.ID)
	if err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, model.ProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		UserID:      data.UserID,
		CategoryID:  data.CategoryID,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	})
}

func (p *productController) Update(ctx *gin.Context) {
	var req model.ProductUpdate

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.service.Update(req); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}

func (p *productController) Delete(ctx *gin.Context) {
	var req model.ProductUri

	if err := ctx.ShouldBindUri(&req); err != nil {
		view.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.service.Delete(req.ID); err != nil {
		view.JSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	view.JSON(ctx, http.StatusOK, nil)
}
