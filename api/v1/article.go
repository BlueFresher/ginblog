package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

// 查询文章是否存在

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章
func GetCategoryArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageSize = -1
	}
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, total := model.GetCategoryArticle(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询某个文章
func GetAeticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	// 查询所有文章时，走redis缓存
	if len(title) == 0 {
		n, _ := model.DBrs.Exists("articles").Result()

		var data []model.Article
		var code int
		var total int
		if n > 0 {
			log.Println("走Redis缓存.")
			timee := model.DBrs.TTL("articles").Val()
			log.Println("缓存过期剩余时间：", timee)
			data, total, code = model.GetArticlesRedist(pageNum, pageSize)
		} else {

			data, code, total = model.GetArticle(pageSize, pageNum)
			// total = len(data)
			_, err := model.DeleteArticlesRedis()
			if err == nil {
				articles := make([]string, 0)
				totalJson, err := json.Marshal(total)
				if err != nil {
					code = errmsg.ERROR_REDIS_DESERIALIZE_WRONG
				}
				articles = append(articles, string(totalJson))
				for _, v := range data {
					article, err := json.Marshal(v)
					if err != nil {
						code = errmsg.ERROR_REDIS_DESERIALIZE_WRONG
						break
					}
					articles = append(articles, string(article))
				}
				model.AddArticlesRedis(articles)
			} else {
				log.Println("Redis缓存清空失败")
			}

		}

		// data, code, total := model.GetArticle(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	// 查询分类文章时
	data, code, total := model.SearchArticle(title, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
