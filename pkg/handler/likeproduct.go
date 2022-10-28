package handler

import (
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLikeProduct(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input perfume.LikeProduct
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.ProductId = int(productId)
	input.UserId = userId

	id, err := h.services.LikeProduct.Create(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getLikeProductById(ctx *gin.Context) {
	likeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	like, err := h.services.LikeProduct.GetById(int(likeId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, like)
}

func (h *Handler) updateLikeProduct(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	likeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input perfume.UpdateLikeProduct
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.LikeProduct.Update(int(likeId), input, userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteLikeProduct(ctx *gin.Context) {
	likeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.LikeProduct.Delete(int(likeId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getAllLikeProductResponse struct {
	Data []perfume.LikeProduct `json:"data"`
}

func (h *Handler) getLikesProductByProductId(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reviews, err := h.services.LikeProduct.GetAll(int(productId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllLikeProductResponse{Data: reviews})
}
