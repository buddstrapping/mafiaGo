package main

import (
	"mafiaGo/controller/day"
	"mafiaGo/controller/dead"
	"mafiaGo/controller/night"
	"mafiaGo/controller/start"
	"mafiaGo/controller/userState"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	userState.LoadOfflineList()
	router := gin.Default()
	//router.LoadHTMLGlob("view/*")
	router.Use(static.Serve("/", static.LocalFile("./view/build", true)))

	r := router.Group("/start")
	{
		r.POST("/setNum", start.SetJobNum)
		r.POST("/reset", start.ResetGame)
		r.POST("/ready", start.Ready)
	}

	d := router.Group("/day")
	{
		d.POST("/nightSwitch", day.GoToNight)
		d.POST("/check", day.GetCareer)
		d.POST("/liveState", day.SetLiveState)
		d.POST("/countdown", day.Countdown)
	}

	n := router.Group("/night")
	{
		n.POST("/setTarget", night.SetTarget)
	}

	de := router.Group("/dead")
	{
		de.POST("/revive", dead.Revive)
		de.POST("/sendmsg", dead.SendChatMessage)
	}

	router.Run(":80")
}
