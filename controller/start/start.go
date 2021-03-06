package start

import (
	"fmt"
	"net/http"

	"mafiaGo/controller/userState"
	"mafiaGo/model"

	"github.com/gin-gonic/gin"
)

/* 메인 페이지 */
func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* 직업 수 설정 */
func SetJobNum(c *gin.Context) {
	var set model.VarSet
	c.BindJSON(&set)
	c.Request.ParseForm()
	res := userState.SetGameRule(set)
	fmt.Printf("[SET] all : %d mafia: %d doc: %d pol: %d %s\n", set.AllNum, set.MafiaNum, set.DocNum, set.PolNum, res)
	c.String(http.StatusOK, res)
}

/* 게임 초기화 */
func ResetGame(c *gin.Context) {
	res := userState.ResetGame()

	c.String(http.StatusOK, res)
}

/* 준비 상태 전환 */
func Ready(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	fmt.Printf("Requested from %s with %s\n", user.Name, user.Career)

	res := userState.UpdateOnline(user)

	if res == 200 || res == 201 {
		c.String(res, "Ready")
	} else {
		c.String(res, "Not Available User")
	}
}
