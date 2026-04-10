package handler

import (
	"server/internal/core/service"
	"server/internal/dto"
	appErr "server/internal/errors"
	"server/internal/mapper"
	"server/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{service: s}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var createMovieRequest dto.CreateMovie

	if err := c.ShouldBindJSON(&createMovieRequest); err != nil {
		response.HandleError(c, appErr.InvalidInput("invalid request body", nil))
		return
	}

	movie, err := h.service.CreateMovie(createMovieRequest)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.OK(c, mapper.ToMovieResponse(movie))

}

func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	idParam := c.Param("id") 

	id, err := uuid.Parse(idParam)
	if err != nil {
		response.HandleError(c,appErr.InvalidInput("invalid movie ID", nil))
		return
	}
	movie, err := h.service.GetMovieByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.OK(c, mapper.ToMovieResponse(movie))
}

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := h.service.GetAllMovies()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.OK(c, mapper.ToMovieResponseList(movies))
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.HandleError(c, appErr.InvalidInput("invalid movie ID", nil))
		return
	}
	var updateMovieRequest dto.UpdateMovie
	if err := c.ShouldBindJSON(&updateMovieRequest); err != nil {
		response.HandleError(c, appErr.InvalidInput("invalid request body", nil))
		return
	}
	movie, err := h.service.UpdateMovie(id, updateMovieRequest)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.OK(c, mapper.ToMovieResponse(movie))
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.HandleError(c, appErr.InvalidInput("invalid movie ID", nil))
		return
	}
	err = h.service.DeleteMovie(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.OK(c, nil)
}