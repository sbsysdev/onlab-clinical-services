package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/OnLab-Clinical/onlab-clinical-services/configs"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth"
	"github.com/OnLab-Clinical/onlab-clinical-services/db"
	"github.com/OnLab-Clinical/onlab-clinical-services/utils"
)

func main() {
	args := os.Args
	// Configure env
	if len(args) > 2 && args[2] == "local" {
		if err := godotenv.Load(".env.app.local"); err != nil {
			panic(err)
		}
	}

	ctx := context.Background()

	// Configure db connection
	connection := configs.ConfigurePostgreSQLConnection(
		utils.GetEnv("DB_HOST", ""),
		utils.GetEnv("DB_USER", ""),
		utils.GetEnv("DB_PASSWORD", ""),
		utils.GetEnv("DB_NAME", ""),
		utils.GetEnv("DB_PORT", ""),
	)

	// Configure migration
	if len(args) > 1 && args[1] == "migrate" {
		db.PublicMigration(connection)
	}

	// TODO: Configure cache

	// TODO: Configure file storage

	// Configure http server
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:           []string{"http://localhost:5173", "http://localhost:4173"},
		AllowMethods:           []string{"GET", "POST", "PUT"},
		AllowHeaders:           []string{"Origin", "Content-Type", "Authorization", "Accept-Language"},
		AllowCredentials:       false,
		ExposeHeaders:          []string{"Content-Length", "content-type"},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
	api := router.Group("/api")

	// Configure modules
	auth.AuthModule{
		Context:                ctx,
		Connection:             connection,
		SubscribeEvent:         configs.SubscribeDomainEvent,
		PublishEvent:           configs.PublishDomainEvent,
		ConfigureEventHandlers: configs.ConfigureEventHandlers,
		Router:                 api.Group("/auth"),
	}.LoadModule()

	router.Run(":8080")

	fmt.Println("OnLab-Clinical Services")
}
