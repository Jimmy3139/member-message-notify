package member

import (
	"grpc-to-do-list/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// MemberService  結構
type MemberService struct {
	Logger *logrus.Logger
	Config *config.Configurations
}

// MemberService 建構
func NewMemberService(l *logrus.Logger, c *config.Configurations) *MemberService {
	return &MemberService{
		Logger: l,
		Config: c}
}

// @Summary 註冊會員
// @Description  註冊會員
// @Version 1.0
// @Accept application/json
// @Param version path number true "version"
// @router /api/v{version}/Device/Register [post]
// @Param Body body RegisterDeviceViewModel true "The body to create a thing"
// @Tags Device
func (s *MemberService) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.Logger.Info("註冊會員")

		c.JSON(200, gin.H{
			"message": "註冊會員",
		})

	}
}
