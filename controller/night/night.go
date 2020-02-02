package night

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 메인 페이지 */
func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* 타겟 셋팅 */
func SetTarget(c *gin.Context) {

}