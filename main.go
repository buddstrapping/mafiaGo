package main

import (
	"mafiaGo/controller/day"
	"mafiaGo/controller/dead"
	"mafiaGo/controller/night"
	"mafiaGo/controller/root"
	"mafiaGo/model"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var offline []model.Account
var online []model.Account
var liveUser []model.Account
var deadUser []model.Account

func main() {
	makeAccount()

	router := gin.Default()
	//router.LoadHTMLGlob("view/*")
	router.Use(static.Serve("/", static.LocalFile("./view/build", true)))

	r := router.Group("/")
	{
		r.GET("/", root.Get)
		r.POST("/setNum", root.SetJobNum)
		r.POST("/ready", root.Ready)
	}

	d := router.Group("/day")
	{
		d.GET("/", day.Get)
		d.POST("/nightSwitch", day.GoToNight)
		d.POST("/career", day.GetCareer)
		d.POST("/liveState", day.SetLiveState)
		d.POST("/countdown", day.Countdown)
	}

	n := router.Group("/night")
	{
		n.GET("/", night.Get)
		n.POST("/setTarget", night.SetTarget)
	}

	de := router.Group("/dead")
	{
		de.GET("/", dead.Get)
		de.POST("/revive", dead.Revive)
		de.POST("/sendmsg", dead.SendChatMessage)
	}

	router.Run(":80")
}

func makeAccount() {
	b, err := ioutil.ReadFile("./model/account_list.json")

	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	json.Unmarshal(b, &offline)
	fmt.Println(offline)
}
