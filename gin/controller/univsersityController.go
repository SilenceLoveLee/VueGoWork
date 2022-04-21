package controller

import (
	"fmt"
	"gin/common"
	"gin/models"
	"gin/pkg/app"
	"gin/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get University
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectUniversity [get]
//主查询页面 (中国大学主榜排名)
func SelectUniversity(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.FirstRunk //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.FirstRunk                 //查询中国大学总榜
		result := DB.Table("FirstRunk").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                   //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByZongBang
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByZongBang [get]
//查询中国大学总榜排名
func SelectByZongBang(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.CountRank //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.CountRank                 //查询中国大学总榜
		result := DB.Table("CountRank").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                   //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")

	}

}

// @Summary Get UniversityByMedicineType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByMedicineType [get]
//查询中国医药类大学排名
func SelectByMedicineType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.MedicineType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.MedicineType                 //查询中国大学总榜
		result := DB.Table("MedicineType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                      //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByFinanceType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByFinanceType [get]
//查询中国财经类大学排名
func SelectByFinanceType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.FinanceType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.FinanceType                 //查询中国大学总榜
		result := DB.Table("FinanceType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                     //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByLanguageType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByLanguageType [get]
//查询语言类大学排名
func SelectByLanguageType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.LanguageType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.LanguageType                 //查询中国大学总榜
		result := DB.Table("LanguageType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                      //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByLawType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByLawType [get]
//查询政法类大学排名
func SelectByLawType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.LawType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.LawType                 //查询中国大学总榜
		result := DB.Table("LawType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                 //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByMinZuType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByMinZuType [get]
//查询民族类大学排名
func SelectByMinZuType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.MinZuType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.MinZuType                 //查询中国大学总榜
		result := DB.Table("MinZuType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                   //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityBySportsType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectBySportsType [get]
//查询体育类大学排名
func SelectBySportsType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.SportsType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.SportsType                 //查询中国大学总榜
		result := DB.Table("SportsType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                    //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByCooperateType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByCooperateType [get]
//查询合作办学大学排名
func SelectByCooperateType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.CooperateType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.CooperateType                 //查询中国大学总榜
		result := DB.Table("CooperateType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                       //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByMinBanType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByMinBanType [get]
//查询民办排名
func SelectByMinBanType(c *gin.Context) {
	DB := common.GetDB()
	var schoolName string
	request := make(map[string]interface{}) //新建一个map类型的变量request，保存前端传来的数据
	//获取参数
	fmt.Println(request)
	c.ShouldBind(&request)                       //将前端传过来的数据绑定到该request上
	schoolName = request["School_name"].(string) //取得相应的map值 放到结构体上
	if schoolName != "" {
		// 根据学校名查询
		var school models.MinBanType //查询中国大学总榜
		result := DB.Where("School_name = ?", schoolName).First(&school)
		c.JSON(http.StatusOK, gin.H{ //多传一些参数 可能会用到
			"code":    "200",
			"schools": school, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	} else {
		// 查询所有
		var schools []models.MinBanType                 //查询中国大学总榜
		result := DB.Table("MinBanType").Find(&schools) //查询所有
		c.JSON(http.StatusOK, gin.H{                    //多传一些参数 可能会用到
			"code":    "200",
			"schools": schools, //试试传一个列表
			"msg":     "成功",
		})
		app.OK(c, result, "OK")
	}
}

// @Summary Get UniversityByArtType
// @Produce  json
// @Param school_name query string false "School_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /selectByArtType [get]
//查询合艺术类大学排名
func SelectByArtType(c *gin.Context) {
	schoolName := ""
	if arg := c.Query("school_name"); arg != "" {
		schoolName = arg
	}
	if schoolName != "" {
		// 根据学校名查询
		err, info := models.ArtTypeByName(schoolName)
		if err != nil {
			app.Error(c, e.ERROR, err, err.Error())
			return
		}
		app.OK(c, info, "OK")
	} else {
		// 查询所有
		err, info := models.ArtTypeSchools()
		if err != nil {
			app.Error(c, e.ERROR, err, err.Error())
			return
		}
		app.OK(c, info, "OK")
	}
}
