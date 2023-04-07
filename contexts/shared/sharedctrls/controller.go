package sharedctrls

import "github.com/gin-gonic/gin"

type Controller interface {
	Handle(*gin.Context)
}
