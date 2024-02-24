package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemService struct {
}

func NewSystemService() *SystemService {
	return &SystemService{}
}

// @Description Home回應 ok，用來確認服務是否已啟動
// @Version 1.0
// @Accept */*
// @Router / [get]
// @Tags Home
func (s *SystemService) GetHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
