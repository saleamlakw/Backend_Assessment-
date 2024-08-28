package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/LoanTracker/internal/tokenutil"
)
func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be 'Bearer <token>'"})
			return
		}
		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, secret)

		if err != nil || !authorized {
		    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := tokenutil.ExtractUserClaimsFromToken(authToken, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set("x-user-id", claims["id"])
		c.Set("x-user-role", claims["role"])
		c.Set("x-user-refresh-data-id", claims["refresh_data_id"])
		c.Next()
	}
}

func IsAdminMiddleware(c *gin.Context){
	role, exists := c.Get("x-user-role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		c.Abort()
		return
	}

	urole, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		c.Abort()
		return
	}


	if urole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
		c.Abort()
		return
	}
	c.Next()
}