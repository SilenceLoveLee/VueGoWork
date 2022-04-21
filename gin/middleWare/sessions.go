package middleWare

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const KEY = "ChinaSoccer"

func EnabledCooikeSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SESSIONID", store)
}
