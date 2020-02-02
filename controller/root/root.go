package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 메인 페이지 */
func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* 직업 수 셋팅 */
func SetJobNum(c *gin.Context) {

}

/* 준비 상태 전환 */
func Ready(c *gin.Context) {

}
