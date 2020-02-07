package night

import (
	"fmt"
	"mafiaGo/controller/userState"
	"mafiaGo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 메인 페이지 */
func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* 타겟 셋팅 */
func SetTarget(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	res := userState.ProcessResult(user)

	c.String(res, "OK")
}

/* CheckRes : 결과 확인 버튼 */
func CheckRes(c *gin.Context) {
	var user model.User
	var liveAll bool
	c.BindJSON(&user)

	res, liveAll := userState.CheckRes(user)

	type restype struct {
		Msg     string `json:"msg"`
		LiveAll bool   `json:"liveAll"`
	}

	m := restype{res, liveAll}
	fmt.Printf("[RESTYPE]: %s\n", res)
	c.JSON(http.StatusOK, m)
}

func DeadRequest(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	res := userState.DeadRequest(user)

	c.String(http.StatusOK, res)
}
