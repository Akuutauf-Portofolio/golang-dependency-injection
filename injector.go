//go:build wireinject
// +build wireinject

package main

import (
	"belajar-go-lang-restful-api/app"
	"belajar-go-lang-restful-api/controller"
	"belajar-go-lang-restful-api/middleware"
	"belajar-go-lang-restful-api/repository"
	"belajar-go-lang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// menyiapkan provider set
var categorySet = wire.NewSet(
	repository.NewCategoryRepository, 
	// kalau ada yang butuh category repository, maka akan mengirimkan *category repository impl
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),

	service.NewCategoryService,
	// kalau ada yang butuh category service, maka akan mengirimkan *category service impl
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),

	controller.NewCategoryController,
	// kalau ada yang butuh category controller, maka akan mengirimkan *category controller impl
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

// membuat injector untuk initialized server
func InitializedServer() *http.Server {
	// db tipe nya struct (tidak perlu binding interface)
	wire.Build(
		app.NewDB,
		wire.Value([]validator.Option{}), // tambahkan agar injector mengerti untuk menangani validator
		validator.New,
		categorySet,
		app.NewRouter,

		// jikalau ada yang butuh handler
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}