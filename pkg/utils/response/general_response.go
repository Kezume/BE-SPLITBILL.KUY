package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"success": false,
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
}

func SuccessWithMeta(c *gin.Context, data any, meta any) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
		"meta":    meta,
	})
}
