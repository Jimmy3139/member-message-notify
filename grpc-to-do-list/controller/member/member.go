package member

import (
	"context"
	"fmt"
	"grpc-to-do-list/config"
	"net/http"

	usecase "grpc-to-do-list/services/member"

	pb "grpc-to-do-list/helpers/grpc/notify"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// MemberService  結構
type MemberService struct {
	Logger        *logrus.Logger
	Config        *config.Configurations
	MemberUseCase *usecase.MemberUseCase
}

// MemberService 建構
func NewMemberService(l *logrus.Logger, c *config.Configurations) *MemberService {
	return &MemberService{
		Logger:        l,
		Config:        c,
		MemberUseCase: usecase.NewMemberService()}
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
		// 呼叫服務
		s.MemberUseCase.Register()

		// 呼叫grpc server
		// 連接 gRPC 用戶服務
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println("could not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewNotificationServiceClient(conn)

		// 調用 gRPC 用戶服務的註冊方法
		response, err := client.SendNotification(context.Background(), &pb.NotificationRequest{
			Message: "註冊會員",
		})
		fmt.Println(response)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(200, gin.H{
			"message": "註冊會員",
		})

	}
}
