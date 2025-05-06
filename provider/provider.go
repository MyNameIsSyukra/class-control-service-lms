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
	kelasRepositoryMember := repository.NewStudentRepository(db)
	kelasService := service.NewKelasService(kelasRepository , kelasRepositoryMember)

	do.Provide(injector, func(i *do.Injector) (controller.KelasController, error) {
		return controller.NewKelasController(kelasService), nil
	})
}

func ProvideMemberDependency(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, "db")

	memberRepository := repository.NewStudentRepository(db)
	memberRepositoryKelas := repository.NewKelasRepository(db)
	memberService := service.NewMemberService(memberRepository, memberRepositoryKelas)

	do.Provide(injector, func(i *do.Injector) (controller.MemberController, error) {
		return controller.NewMemberController(memberService), nil
	})
}

func RegisterProviders(injector *do.Injector) {
	InitDatabase(injector)
	ProvideKelasDependency(injector)
	ProvideMemberDependency(injector)
}

