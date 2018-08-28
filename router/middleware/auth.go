package middleware

import (
	"encoding/base64"

	"github.com/blog-web/common/components"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		data, err := base64.URLEncoding.DecodeString(token)
		if err != nil {
			panic(err)
		}

		token = string(data)

		idStr := ""
		if token != "" {
			if flag, userId := components.RequireTokenAuthentication(token); flag == true {
				idStr = userId
			}
		}
		c.Set("userId", idStr)
		c.Next()
	}
}
