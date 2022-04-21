package routers

import (
	middlewares "gin/middleWare"
	_ "gin/models"
	"github.com/gin-gonic/gin"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	SysBaseRouter(r)
	CollectRoute(r)
}

func sysCheckRoleRouter(r *gin.RouterGroup, authMiddle gin.HandlerFunc) {
	SysBaseRouterWithMiddleWare(r, authMiddle)
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())
	r.Use(middlewares.EnabledCooikeSession())

	g := r.Group("")
	//jwt := middleWare.JWT()

	sysNoCheckRoleRouter(g)
	//sysCheckRoleRouter(g,jwt)
	return r
}
