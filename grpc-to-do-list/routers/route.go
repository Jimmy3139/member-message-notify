package route

import (
	"grpc-to-do-list/config"
	system "grpc-to-do-list/controller/health"
	"grpc-to-do-list/controller/member"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Route struct {
	*gin.RouterGroup
	memberService *member.MemberService
	systemService *system.SystemService
}

// NewApiRoute 建構子
func NewApiRoute(group *gin.RouterGroup, l *logrus.Logger, conf *config.Configurations) *Route {
	return &Route{
		group,
		member.NewMemberService(l, conf),
		system.NewSystemService(),
	}
}

func (route *Route) RouteApi() {
	//健康檢查
	route.GET("/health", route.systemService.GetHealthCheck())
	//註冊會員
	route.POST("/member/register", route.memberService.Register())

}
