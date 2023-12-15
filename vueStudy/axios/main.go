package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

type article struct {
	title string
	category string
	time string
	state string
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/getAll", getAll)
	r.GET("/search", search)
	r.POST("/add", add)

	r.Run()
}

func getAll(ctx *gin.Context) {
	articleList := make(map[int]*article)
	articleList[0] = &article{
		title: "医疗反腐绝非砍医护收入",
		category: "时事",
		time: time.UnixDate,
		state: "已经发布",
	}
	articleList[1] = &article{
		title: "医疗反腐绝非砍医护收入",
		category: "时事",
		time: time.UnixDate,
		state: "已经发布",
	}
	articleList[2] = &article{
		title: "医疗反腐绝非砍医护收入",
		category: "时事",
		time: time.UnixDate,
		state: "已经发布",
	}
	ctx.JSON(200, gin.H {
		"1": "1",
		"2": "2", 
		"3": "3",
	})
}

func search(ctx *gin.Context) {

}

func add(ctx *gin.Context) {

}
