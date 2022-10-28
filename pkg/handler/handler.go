package handler

import (
	"github.com/Ig0rVItalevich/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		products := api.Group("/products")
		{
			products.GET("/", h.getAllProducts)
			products.POST("/", h.createProduct)
			products.GET("/:id", h.getProductById)
			products.PUT("/:id", h.updateProduct)
			products.DELETE("/:id", h.deleteProduct)

			reviews := products.Group(":id/reviews")
			{
				reviews.POST("/", h.createReview)
				reviews.GET("/", h.getReviewsByProductId)
			}

			likeProducts := products.Group(":id/like_products")
			{
				likeProducts.POST("/", h.createLikeProduct)
				likeProducts.GET("/", h.getLikesProductByProductId)
			}
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("/:id", h.getReviewById)
			reviews.PUT("/:id", h.updateReview)
			reviews.DELETE("/:id", h.deleteReview)

			likeReviews := reviews.Group(":id/like_reviews")
			{
				likeReviews.POST("/", h.createLikeReview)
				likeReviews.GET("/", h.getLikesReviewByReviewId)
			}
		}

		likeProducts := api.Group("/like_products")
		{
			likeProducts.GET("/:id", h.getLikeProductById)
			likeProducts.PUT("/:id", h.updateLikeProduct)
			likeProducts.DELETE("/:id", h.deleteLikeProduct)
		}

		likeReviews := api.Group("/like_reviews")
		{
			likeReviews.GET("/:id", h.getLikeReviewById)
			likeReviews.PUT("/:id", h.updateLikeReview)
			likeReviews.DELETE("/:id", h.deleteLikeReview)
		}

		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrder)
			orders.GET("/", h.getOrdersByUserId)
			orders.GET("/:id", h.getOrderById)
			orders.PUT("/:id", h.updateOrder)
			orders.DELETE("/:id", h.deleteOrder)

			products := orders.Group(":id/products")
			{
				products.GET("/", h.getProductsByOrder)
			}
		}
	}

	return router
}
