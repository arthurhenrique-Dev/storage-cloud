package startup

import (
	"database/sql"
	"fmt"
	"net/http"

	"storage-api/core/configuration"
	"storage-api/core/database/connection"
	"storage-api/core/database/migrations"
	itemsHandler "storage-api/modules/items/handlers"
	itemsRepo "storage-api/modules/items/repositories"
	itemsService "storage-api/modules/items/services"
	suppliersHandler "storage-api/modules/suppliers/handlers"
	suppliersRepo "storage-api/modules/suppliers/repositories"
	suppliersService "storage-api/modules/suppliers/services"
	usersHandler "storage-api/modules/users/handlers"
	usersRepo "storage-api/modules/users/repositories"
	usersService "storage-api/modules/users/services"
	"storage-api/shared/operations"
)

func Startup() error {
	var (
		cfg *configuration.Config
		db  *sql.DB
		mux *http.ServeMux
	)

	return operations.NewPipeline("Application Startup").
		Step("Load Configuration", func() error {
			cfg = configuration.Load()
			return nil
		}).
		Step("Connect Database", func() error {
			var err error
			db, err = connection.Connect(cfg.DatabaseURL)
			return err
		}).
		Step("Run Database Migrations", func() error {
			return migrations.Run(db)
		}).
		Step("Initialize Handlers and Routes", func() error {
			mux = http.NewServeMux()

			itemRepository := itemsRepo.NewItemRepository(db)
			itemService := itemsService.NewItemService(itemRepository)
			itemHandler := itemsHandler.NewItemHandler(itemService)
			itemHandler.RegisterRoutes(mux)

			supplierRepository := suppliersRepo.NewSupplierRepository(db)
			supplierService := suppliersService.NewSupplierService(supplierRepository)
			supplierHandler := suppliersHandler.NewSupplierHandler(supplierService)
			supplierHandler.RegisterRoutes(mux)

			userRepository := usersRepo.NewUserRepository(db)
			userService := usersService.NewUserService(userRepository)
			userHandler := usersHandler.NewUserHandler(userService)
			userHandler.RegisterRoutes(mux)

			return nil
		}).
		Step("Start HTTP Server", func() error {
			fmt.Printf("HTTP Server listening on port %s\n", cfg.Port)
			return http.ListenAndServe(cfg.Port, mux)
		}).
		Execute()
}
