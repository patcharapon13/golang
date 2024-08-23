package main

import (
	"clean_architecture/adapters"
	"clean_architecture/entities"
	"clean_architecture/usecases"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 5432
	usr    = "dbusr"
	pass   = "dbpass#123"
	dbname = "progessDB"
)

func main() {
	app := fiber.New()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", host, port, usr, pass, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail connect to database")
	}

	err = db.AutoMigrate(&entities.Order{})
	if err != nil {
		panic("Migration run has been failed")
	}

	println(db)
	println("Connect is successful")

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	err = app.Listen(":8080")
	if err != nil {
		panic("fail to start application")
	}

}
