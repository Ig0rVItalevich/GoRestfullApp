package handler

import (
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createReview(ctx *gin.Context) {
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

	var input perfume.Review
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.ProductId = int(productId)
	input.UserId = userId

	id, err := h.services.Review.Create(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getReviewById(ctx *gin.Context) {
	reviewId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	review, err := h.services.Review.GetById(int(reviewId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, review)
}

func (h *Handler) updateReview(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reviewId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input perfume.UpdateReview
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Review.Update(int(reviewId), input, userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteReview(ctx *gin.Context) {
	reviewId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Review.Delete(int(reviewId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getAllReviewResponse struct {
	Data []perfume.Review `json:"data"`
}

func (h *Handler) getReviewsByProductId(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reviews, err := h.services.Review.GetAll(int(productId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllReviewResponse{Data: reviews})
}
