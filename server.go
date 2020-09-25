package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db "touristapp.com/db"
	routers "touristapp.com/routers"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	if err := db.Connect(); err != nil {
		log.Fatal("DB connection failed")
	} else {
		log.Println("DB Connection successfull!")
	}

	routers.Routes(app)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Exiting, Bye!")
		db.DB.Close()
		fmt.Println("DB connection closed")
		os.Exit(0)
	}()

	log.Fatal(app.Listen(":7000"))
}
