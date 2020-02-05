package day

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

/* 밤 전환 스위치 */
func GoToNight(c *gin.Context) {

}

/* 직업 확인 */
func GetCareer(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	career := userState.CheckCareer(user)

	fmt.Printf("[Day] %s %s->%s", user.Name, user.Career, career)
	c.String(http.StatusOK, career)
}

/* 직업 테스팅 */
func TestCareer(c *gin.Context) {

}

/* 사망 신고 */
func SetLiveState(c *gin.Context) {

}

/* 회의 시간 카운트 다운 */
func Countdown(c *gin.Context) {

}
