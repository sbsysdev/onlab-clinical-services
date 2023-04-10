package main

import (
	"context"
	"fmt"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	// TODO: Configure db

	// TODO: Configure cache

	// TODO: Configure file storage

	// Configure http server
	router := gin.Default()
	api := router.Group("/api")

	// Configure modules
	auth.AuthModule{
		Context: ctx,
		// Connection: ,
		Router: api.Group("/auth"),
	}.LoadModule()

	router.Run(":8080")

	fmt.Println("OnLab-Clinical Services")
}
