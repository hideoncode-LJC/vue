package main

import (
	"axios/model"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/getAll", getAll)
	r.GET("/search", search)
	r.POST("/add", add)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func getAll(ctx *gin.Context) {
	db := model.GetDatabaseConnection()

	var articles []model.Article

	result := db.Find(&articles)

	fmt.Println("查询到的记录个数为", result.RowsAffected)

	data := make(map[string]any)

	for i := 0; i < len(articles); i++ {
		data[fmt.Sprint(i)] = articles[i]
	}

	ctx.JSON(200, data)
}

func search(ctx *gin.Context) {

}

func add(ctx *gin.Context) {
	db := model.GetDatabaseConnection()

	//var article model.Article
	//
	//if err := ctx.ShouldBindJSON(&article); err != nil {
	//	ctx.JSON(400, gin.H{
	//		"message": "添加失败",
	//	})
	//} else {
	//	article.Time = time.Now()
	//	db.Create(&article)
	//	ctx.JSON(200, gin.H{
	//		"message": "添加成功",
	//	})
	//}

	db.Create(&model.Article{
		Title:    ctx.PostForm("title"),
		Category: ctx.PostForm("category"),
		Time:     time.Now(),
		State:    ctx.PostForm("state"),
	})

	ctx.JSON(200, gin.H{
		"message": "添加成功",
	})

}
