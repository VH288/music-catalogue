package memberships

import (
	"github.com/VH288/music-catalogue/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(request memberships.SignUpRequest) error
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
	route := h.Group("/memberships")
	route.POST("/signup", h.SignUp)
}
