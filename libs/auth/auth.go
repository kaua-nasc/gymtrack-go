package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type UserContextKeyType string

const UserContextKey UserContextKeyType = "user"
const TokenContextKey UserContextKeyType = "token"

type UserType string

const (
	Trainer UserType = "PERSONAL_TRAINER"
	Client  UserType = "CLIENT"
)

type AuthUser struct {
	ID   string
	Type UserType
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

		typedUserType := UserType(userType)

		if typedUserType != Client && typedUserType != Trainer {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type"})
			c.Abort()
			return
		}

		c.Set(string(UserContextKey), AuthUser{
			ID:   userID,
			Type: typedUserType,
		})
		c.Set(string(TokenContextKey), tokenString)

		// Inject into request context so it propagates to service layer
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, string(UserContextKey), AuthUser{
			ID:   userID,
			Type: typedUserType,
		})
		ctx = context.WithValue(ctx, string(TokenContextKey), tokenString)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// RolesMiddleware ensures the user has one of the required roles
func RolesMiddleware(allowedTypes ...UserType) gin.HandlerFunc {
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

func GetAuthUser(ctx *gin.Context) (AuthUser, bool) {
	user, ok := ctx.Value(string(UserContextKey)).(AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("unauthorized"))
		return AuthUser{}, false
	}

	return user, true
}

func GetPagination(ctx *gin.Context) (string, int) {
	cursor := ctx.Query("cursor")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit <= 0 {
		limit = 20
	}
	return cursor, limit
}
