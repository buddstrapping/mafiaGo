package dead

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 메인 페이지 */
func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* 부활 */
func Revive(c *gin.Context) {

}

/* 채팅 메시지 전송 */
func SendChatMessage(c *gin.Context) {

}
