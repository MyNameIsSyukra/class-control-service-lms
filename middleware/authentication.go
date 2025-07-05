package middleware

import (
	"LMSGo/dto"
	"LMSGo/service"
	"LMSGo/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		
		if authHeader == "" {
			response := utils.FailedResponse(dto.MESSAGE_FAILED_TOKEN_NOT_FOUND)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.FailedResponse(dto.MESSAGE_FAILED_TOKEN_NOT_FOUND)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.FailedResponse(dto.MESSAGE_FAILED_TOKEN_NOT_VALID)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := utils.FailedResponse(dto.MESSAGE_FAILED_TOKEN_NOT_VALID)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.FailedResponse(dto.MESSAGE_FAILED_TOKEN_NOT_VALID)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("token", authHeader)
		ctx.Set("uuid", userId)
		ctx.Next()
	}
}

func RequireTeacherRole(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil token dari context yang sudah di-set oleh middleware Authenticate
		token, exists := ctx.Get("token")
		if !exists {
			response := utils.FailedResponse("Token not found in context")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString, ok := token.(string)
		if !ok {
			response := utils.FailedResponse("Invalid token format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Dapatkan role dari token
		role, err := jwtService.GeRoleByToken(tokenString)
		if err != nil {
			response := utils.FailedResponse("Failed to get user role")
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		// Periksa apakah role adalah teacher
		if role != "teacher" && role != "admin" {
			response := utils.FailedResponse("Access denied: Teacher role required")
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// Set role ke context untuk digunakan di handler selanjutnya
		ctx.Set("role", role)
		ctx.Next()
	}
}

func RequireAdminRole(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil token dari context yang sudah di-set oleh middleware Authenticate
		token, exists := ctx.Get("token")
		if !exists {
			response := utils.FailedResponse("Token not found in context")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString, ok := token.(string)
		if !ok {
			response := utils.FailedResponse("Invalid token format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Dapatkan role dari token
		role, err := jwtService.GeRoleByToken(tokenString)
		if err != nil {
			response := utils.FailedResponse("Failed to get user role")
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		// Periksa apakah role adalah teacher
		if role != "admin" {
			response := utils.FailedResponse("Access denied: Teacher role required")
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// Set role ke context untuk digunakan di handler selanjutnya
		ctx.Set("role", role)
		ctx.Next()
	}
}