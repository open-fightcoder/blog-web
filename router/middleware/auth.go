package middleware

import (
	"encoding/base64"
	"net/http"

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

		if token != "" {
			if flag, userId := components.RequireTokenAuthentication(token); flag == true {
				c.Set("userId", userId)
				c.Next()
				return
			}
		}
		//c.AbortWithStatusJSON(http.StatusForbidden, base.Fail("auth failure."))
		c.AbortWithStatusJSON(http.StatusForbidden, "")
	}
}
