package provider

import (
	config "LMSGo/config"
	controller "LMSGo/controller"
	repository "LMSGo/repository"
	service "LMSGo/service"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, "db", func (i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func ProvideKelasDependency(injector *do.Injector){
	db := do.MustInvokeNamed[*gorm.DB](injector, "db")

	kelasRepository := repository.NewKelasRepository(db)
	kelasService := service.NewKelasService(kelasRepository)

	do.Provide(injector, func(i *do.Injector) (controller.KelasController, error) {
		return controller.NewKelasController(kelasService), nil
	})
}

func RegisterProviders(injector *do.Injector) {
	InitDatabase(injector)
	ProvideKelasDependency(injector)
}

