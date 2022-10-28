package handler

import (
	"github.com/gin-gonic/gin"
)

type orderInputWithProducts struct {
	Status   string
	Comment  string
	Products []int ``
}

func (h *Handler) createOrder(ctx *gin.Context) {

}

func (h *Handler) getOrderById(ctx *gin.Context) {

}

func (h *Handler) updateOrder(ctx *gin.Context) {

}

func (h *Handler) deleteOrder(ctx *gin.Context) {

}

func (h *Handler) getOrdersByUserId(ctx *gin.Context) {

}

func (h *Handler) getProductsByOrder(ctx *gin.Context) {

}
