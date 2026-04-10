package route

import (
	"server/internal/core/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(r *gin.RouterGroup, h *handler.MovieHandler, jwtSecret string) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("", h.GetAllMovies)
	r.GET("/:id", h.GetMovieByID)
	r.POST("", middleware.JWTAuth(jwtSecret), h.CreateMovie)
	r.PUT("/:id", middleware.JWTAuth(jwtSecret),h.UpdateMovie)
	r.DELETE("/:id", middleware.JWTAuth(jwtSecret),h.DeleteMovie)
}