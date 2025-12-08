package tracks

import (
	"context"

	"github.com/VH288/music-catalogue/internal/middleware"
	"github.com/VH288/music-catalogue/internal/models/spotify"
	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error)
	UpsertTrackActivites(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error
}

type (
	Handler struct {
		*gin.Engine
		service service
	}
)

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		Engine:  api,
		service: service,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/tracks")
	route.Use(middleware.AuthMiddleware())
	route.GET("/search", h.Seacrh)
	route.POST("/liking", h.UpsertTrackActivities)
}
