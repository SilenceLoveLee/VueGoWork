package main

import (
	"fmt"
	"gin/common"
	_ "gin/docs"
	"gin/models"
	"gin/pkg/gredis"
	"gin/pkg/setting"
	"gin/routers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
}

func main() {
	db := common.InitDB() //初始化数据库，这样后边的函数再调用数据库的时候直接获取DB就可以
	fmt.Printf("%T", db)
	gin.SetMode(setting.ServerSetting.RunMode)
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routersInit.Run(endPoint)

	//r = CollectRoute(r)
	//链接数据库

	/*db := common.InitDB()    //初始化数据库，这样后边的函数再调用数据库的时候直接获取DB就可以
	r := gin.Default()       //启动引擎
	r.Use(middleWare.Cors()) //跨域!!!!!!!!
	r.POST("selectUniversity", controller.SelectUniversity)             //主查询页面 (中国大学主榜排名)
	r.POST("selectByZongBang", controller.SelectByZongBang)             //查询中国大学总榜排名
	r.POST("selectUniversityByName", controller.SelectUniversityByName) //根据学校名查询大学返回信息
	r.POST("selectByMedicineType", controller.SelectByMedicineType)     //查询中国医药类大学排名
	r.POST("selectByFinanceType", controller.SelectByFinanceType)       //查询中国财经类大学排名
	r.POST("selectByLanguageType", controller.SelectByLanguageType)     //查询语言类大学排名
	r.POST("selectByLawType", controller.SelectByLawType)               //查询政法类大学排名
	r.POST("selectByMinZuType", controller.SelectByMinZuType)           //查询民族类大学排名
	r.POST("selectBySportsType", controller.SelectBySportsType)         //查询体育类大学排名
	r.POST("selectByCooperateType", controller.SelectByCooperateType)   //查询合作办学大学排名
	r.POST("selectByMinBanType", controller.SelectByMinBanType)         //查询民办排名
	r.POST("selectByArtType", controller.SelectByArtType)               //查询合艺术类大学排名


	//defer db.Close()
	 //无奈之举*/
}
