package main

import (
	"log"

	"github.com/kryast/Crud-5.git/database"
	"github.com/kryast/Crud-5.git/handlers"
	"github.com/kryast/Crud-5.git/models"
	"github.com/kryast/Crud-5.git/repositories"
	"github.com/kryast/Crud-5.git/router"
	"github.com/kryast/Crud-5.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal konek DB:", err)
	}

	db.AutoMigrate(
		&models.Customer{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	orderItemRepo := repositories.NewOrderItemRepository(db)
	orderItemService := services.NewOrderItemService(orderItemRepo)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemService)

	r := router.SetupRouter(orderItemHandler)
	r.Run(":8080")
}
