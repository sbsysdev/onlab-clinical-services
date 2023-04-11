package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/configs"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth"
	"github.com/OnLab-Clinical/onlab-clinical-services/db"
	"github.com/OnLab-Clinical/onlab-clinical-services/utils"
)

func main() {
	ctx := context.Background()

	// Configure db connection
	connection := configs.ConfigurePostgreSQLConnection(
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "user"),
		utils.GetEnv("DB_PASSWORD", "1234"),
		utils.GetEnv("DB_NAME", "onlab_clinical"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	// Configure migration
	if os.Args[1] == "migrate" {
		db.PublicMigration(connection)
	}

	// TODO: Configure cache

	// TODO: Configure file storage

	// Configure http server
	router := gin.Default()
	api := router.Group("/api")

	// Configure modules
	auth.AuthModule{
		Context:    ctx,
		Connection: connection,
		Router:     api.Group("/auth"),
	}.LoadModule()

	router.Run(":8080")

	fmt.Println("OnLab-Clinical Services")
}
