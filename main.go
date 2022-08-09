package main

import (
	"log"
	"os"

	"github.com/dwiiipoetra/restapi-miniproject/service"
	"github.com/dwiiipoetra/restapi-miniproject/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load database")
	}

	// err = models.MigrateMotorcycles(db)
	// if err != nil {
	// 	log.Fatal("could not migrate db")
	// }

	// err = models.MigrateUsers(db)
	// if err != nil {
	// 	log.Fatal("could not migrate db")
	// }

	r := &service.Repository{
		DB: db,
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	r.SetupRoutes(app)

	app.Listen(":3000")
}
