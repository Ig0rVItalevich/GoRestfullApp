package handler

import (
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLikeReview(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reviewtId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input perfume.LikeReview
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	input.ReviewId = int(reviewtId)
	input.UserId = userId

	id, err := h.services.LikeReview.Create(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getLikeReviewById(ctx *gin.Context) {
	likeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	like, err := h.services.LikeReview.GetById(int(likeId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, like)
}

func (h *Handler) updateLikeReview(ctx *gin.Context) {
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

	var input perfume.UpdateLikeReview
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.LikeReview.Update(int(likeId), input, userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteLikeReview(ctx *gin.Context) {
	likeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.LikeReview.Delete(int(likeId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getAllLikeReviewResponse struct {
	Data []perfume.LikeReview `json:"data"`
}

func (h *Handler) getLikesReviewByReviewId(ctx *gin.Context) {
	reviewId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reviews, err := h.services.LikeReview.GetAll(int(reviewId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllLikeReviewResponse{Data: reviews})
}
