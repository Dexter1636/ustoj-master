package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ustoj-master/vo"
)

// Recovery middleware recovers from any panics and writes a 500 if there was one.
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, vo.ResponseMeta{Code: vo.UnknownError})
			log.Println("========================================")
			log.Printf("!!PANIC!! ERR: %s\n", err)
			log.Println("========================================")
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
