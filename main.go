package main

import (
	database "LMSGo/config"
	middleware "LMSGo/middleware"
	migration "LMSGo/migration"
	provider "LMSGo/provider"
	routes "LMSGo/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func args(db *gorm.DB) bool {
    if len(os.Args) > 1 {
        if (os.Args[1] == "migrate") {
            print("Migration Success")
			err := migration.Migrate(db)
			if err != nil {
				log.Fatalf("error running migration: %v", err)
			}
            return false
        }
		if (os.Args[1] == "seed") {
			print("Seeding Success")
			err := migration.Seeder()
			if err != nil {
				log.Fatalf("error running seeder: %v", err)
			}
			return false
		}
		if (os.Args[1] == "rollback") {
			err := migration.Rollback(db)
			if err != nil {
				log.Fatalf("error running rollback: %v", err)
			}
			print("Rollback Success")
			return false
		}
    }
    return true
}


func run(server *gin.Engine) {
	server.Static("/assets", "./assets")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
    var (
		injector = do.New()
	)
    provider.RegisterProviders(injector)
    db := database.SetUpDatabaseConnection()
    defer database.CloseDatabaseConnection(db)

    if !args(db) {
		return
	}
	err := migration.Migrate(db)
	if err != nil {
		log.Fatalf("error running migration: %v", err)
	}
    server := gin.Default()
	server.Use(middleware.CORSMiddleware())
    
	// routes
	routes.RegisterRoutes(server, injector)

	run(server)
    
}