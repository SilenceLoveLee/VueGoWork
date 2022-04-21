package routers

import (
	"gin/controller"
	"github.com/gin-gonic/gin"
)

//func CollectRoute(r *gin.Engine) *gin.Engine {

func CollectRoute(r *gin.RouterGroup) {
	r.GET("/selectUniversity", controller.SelectUniversity)           //主查询页面 (中国大学主榜排名)
	r.GET("/selectByZongBang", controller.SelectByZongBang)           //查询中国大学总榜排名
	r.GET("/selectByMedicineType", controller.SelectByMedicineType)   //查询中国医药类大学排名
	r.GET("/selectByFinanceType", controller.SelectByFinanceType)     //查询中国财经类大学排名
	r.GET("/selectByLanguageType", controller.SelectByLanguageType)   //查询语言类大学排名
	r.GET("/selectByLawType", controller.SelectByLawType)             //查询政法类大学排名
	r.GET("/selectByMinZuType", controller.SelectByMinZuType)         //查询民族类大学排名
	r.GET("/selectBySportsType", controller.SelectBySportsType)       //查询体育类大学排名
	r.GET("/selectByCooperateType", controller.SelectByCooperateType) //查询合作办学大学排名
	r.GET("/selectByMinBanType", controller.SelectByMinBanType)       //查询民办排名
	r.GET("/selectByArtType", controller.SelectByArtType)             //查询合艺术类大学排名
	//return r
}
