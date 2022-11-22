package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
func GetCategoryArticle(cid int, pageSize, pageNum int) ([]Article, int, int) {
	var catArticleList []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", cid).Find(&catArticleList).Error
	db.Model(&catArticleList).Where("cid =?", cid).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CAR_NOT_EXIST, 0
	}
	return catArticleList, errmsg.SUCCESS, total
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id=?", id).First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询文章列表
func GetArticle(title string, pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	if title == "" {
		err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
		// 这是记录总数的，所以total=len(articleList)是不对的
		db.Model(&articleList).Count(&total)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
		return articleList, errmsg.SUCCESS, total
	}
	err = db.Preload("Category").Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
