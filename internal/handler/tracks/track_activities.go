package tracks

import (
	"net/http"

	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var req trackactivities.TrackActivityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")

	err := h.service.UpsertTrackActivites(ctx, userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
