package handler

import (
	"api-gateway/pkg/client"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/suyundykovv/protos/gen/go/statistics/v1"
)

type StatisticsHandler struct {
	statisticsClient *client.StatisticsClient
}

func NewStatisticsHandler(statisticsClient *client.StatisticsClient) *StatisticsHandler {
	return &StatisticsHandler{
		statisticsClient: statisticsClient,
	}
}

func (h *StatisticsHandler) GetUserStatistics(c *gin.Context) {
	timePeriod := c.Query("period")

	if timePeriod == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time Perod is required"})
		return
	}

	req := &pb.GetUserStatisticsRequest{TimePeriod: &timePeriod}
	statistics, err := h.statisticsClient.GetUserStatistics(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving statistics: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}
func (h *StatisticsHandler) GetUserOrdersStatistics(c *gin.Context) {
	userID := c.Param("id")         // <-- берём из пути
	timePeriod := c.Param("period") // <-- берём из пути

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "USer ID is required"})
		return
	}
	if timePeriod == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time period is required"})
		return
	}
	req := &pb.GetUserOrdersStatisticsRequest{
		UserId:     userID,
		TimePeriod: &timePeriod,
	}
	statistics, err := h.statisticsClient.GetUserOrdersStatistics(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving stats: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}
