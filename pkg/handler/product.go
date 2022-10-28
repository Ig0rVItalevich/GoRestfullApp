package handler

import (
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllProductResponse struct {
	Data []perfume.Product `json:"data"`
}

func (h *Handler) getAllProducts(ctx *gin.Context) {
	products, err := h.services.Product.GetAll()
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllProductResponse{Data: products})
}

func (h *Handler) createProduct(ctx *gin.Context) {
	var input perfume.Product
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Product.Create(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.Product.GetById(int(id))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Handler) updateProduct(ctx *gin.Context) {
	var input perfume.UpdateProduct
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Product.Update(int(id), input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Product.Delete(int(id))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
