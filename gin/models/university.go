package models

//2022主榜大学
type FirstRunk struct {
	//gorm.Model
	Rank_id             string `gorm:"column:Rank_id"`             //2022排名
	School_name         string `gorm:"column:School_name"`         //学校名称
	Provinces_cities    string `gorm:"column:Provinces_cities"`    //省市
	School_type         string `gorm:"column:School_type"`         //学校类型
	Total_score         string `gorm:"column:Total_score"`         //总得分
	School_level        string `gorm:"column:School_level"`        //办学层次得分
	Subject_level       string `gorm:"column:Subject_level"`       //学科水平得分
	School_resources    string `gorm:"column:School_resources"`    //办学资源得分
	S_teacherscale      string `gorm:"column:S_teacherscale"`      //师资规模与结构得分
	Talent_Training     string `gorm:"column:Talent_Training"`     //人才培养得分
	Scientific_research string `gorm:"column:Scientific_research"` //科学研究得分
	Service_Society     string `gorm:"column:Service_Society"`     //服务社会得分
	High_talent         string `gorm:"column:High_talent"`         //高端人才得分
	S_majorprojects     string `gorm:"column:S_majorprojects"`     //重大项目与成果得分
	InternationalC      string `gorm:"column:InternationalC"`      //国际竞争力得分
}

//2022总榜大学
type CountRank struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022艺术类名单
type ArtType struct {
	//gorm.Model
	School_name         string `gorm:"column:School_name"`         //学校名称
	Provinces_cities    string `gorm:"column:Provinces_cities"`    //省市
	School_type         string `gorm:"column:School_type"`         //学校类型
	DoctorP             string `gorm:"column:DoctorP"`             //博士点
	MasterP             string `gorm:"column:MasterP"`             //硕士点
	Certification_major string `gorm:"column:Certification_major"` //国家级与认证专业
	Employment_rate     string `gorm:"column:Employment_rate"`     //本科毕业生就业率
	Furtherstudy_rate   string `gorm:"column:Furtherstudy_rate"`   //本科毕业生深造率
}

//2022合作办学大学
type CooperateType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022体育类
type SportsType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022民族类
type MinZuType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022政法类
type LawType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022语言类
type LanguageType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022民办类
type MinBanType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022财经类
type FinanceType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

//2022医药类
type MedicineType struct {
	//gorm.Model
	Rank_id          string `gorm:"column:Rank_id"`          //2022排名
	National_ranking string `gorm:"column:National_ranking"` //全国参考排名
	School_name      string `gorm:"column:Provinces_cities"` //学校名称
	Provinces_cities string `gorm:"column:Provinces_cities"` //省市
	School_type      string `gorm:"column:School_type"`      //学校类型
	Total_score      string `gorm:"column:Total_score"`      //总得分
}

func ArtTypeSchools() (error, []ArtType) {
	var schoolData []ArtType
	//err := db.Table("student").Select("id, stu_id, stu_name, sex, birthday, department, class, is_del, update_time").Where("is_del = false").Find(&studentdata).Error
	err := db.Table("ArtType").Find(&schoolData).Error
	return err, schoolData
}

func ArtTypeByName(school_name string) (error, ArtType) {
	var schoolData ArtType
	//err := db.Table("student").Select("id, stu_id, stu_name, sex, birthday, department, class, is_del, update_time").Where("stu_id = ? and is_del = false", stu_id).Find(&studentdata).Error
	err := db.Table("ArtType").Where("\"School_name\" = ?", school_name).Find(&schoolData).Error
	return err, schoolData
}
