package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

func CheckTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		authParts := strings.Split(authorization, "Bearer ")

		if len(authParts) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": shareddomain.ERRORS_UNAUTHORIZED,
			})

			return
		}

		token := authParts[1]

		issuer, subject, expiration, err := authdomain.DecodeToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": err.Error(),
			})

			return
		}

		if issuer != "OnLab-Clinical" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": authdomain.ERRORS_TOKEN_UNKNOWN,
			})

			return
		}

		if expiration.UTC().Before(time.Now().UTC()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": authdomain.ERRORS_TOKEN_EXPIRED,
			})

			return
		}

		ctx.Set("patientId", subject)

		ctx.Next()
	}
}
