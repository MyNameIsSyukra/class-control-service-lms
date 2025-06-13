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