package main

import (
	database "LMSGo/config"
	migration "LMSGo/migration"
	repository "LMSGo/repository"
	routes "LMSGo/router"
	service "LMSGo/service"
	"os"

	"gorm.io/gorm"
)

func args(db *gorm.DB) bool {
    if len(os.Args) > 1 {
        if (os.Args[1] == "migrate") {
            print("argadasds")
            migration.Migrate(db)
            return false
        }
    }
        return true
}


func main() {
    db := database.SetUpDatabaseConnection()
    defer database.CloseDatabaseConnection(db)

    if !args(db) {
		return
	}

    kelasDB := repository.NewKelasRepository(db)
    kelasUC := service.NewKelasService(kelasDB)
    router := routes.SetupRouter(kelasUC)

    router.Run(":8080")

    
}