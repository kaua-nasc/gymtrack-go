package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserContextKeyType string

const UserContextKey UserContextKeyType = "user"

type AuthUser struct {
	ID   string
	Type string
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Step 1: Extraction (Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default_secret" // Must be identical to NestJS JWT_SECRET
		}

		// Step 2 & 3: Signature Verification (HS256) and Expiration check
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the algorithm is HS256 as per NestJS configuration
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token: " + err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Step 4: Extract Payload Pattern { sub: string, type: string }
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Compatibility with NestJS payload structure
		userID, _ := claims["sub"].(string)
		userType, _ := claims["type"].(string)

		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing subject (sub)"})
			c.Abort()
			return
		}

		c.Set(string(UserContextKey), AuthUser{
			ID:   userID,
			Type: userType,
		})

		c.Next()
	}
}

// RolesMiddleware ensures the user has one of the required roles
func RolesMiddleware(allowedTypes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value(string(UserContextKey)).(AuthUser)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		isAllowed := false
		for _, t := range allowedTypes {
			if user.Type == t {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
