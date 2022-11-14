package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Category struct {
	// gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategory(name string) int {
	var cat Category
	db.Select("id").Where("name=?", name).First(&cat)
	if cat.ID > 0 {
		return errmsg.ERROR_CATNAME_USE
	}
	return errmsg.SUCCESS
}

func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cat Category
	db.Where("id=?", id).First(&cat)
	return cat, errmsg.SUCCESS
}

//　TODO: ＂查询分类下的所有文章

func GetCat(pageSize int, pageNum int) ([]Category, int) {
	var cat []Category
	var total int
	err = db.Find(&cat).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cat, total
}

func EditCat(id int, data *Category) int {
	var cat Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cat).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCat(id int) int {
	var cat Category
	err = db.Where("id=?", id).Delete(&cat).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
