package main

import (
	"github.com/chizidotdev/nuntius/config"
	"github.com/chizidotdev/nuntius/internal/app/db"
	"github.com/chizidotdev/nuntius/internal/app/drivers"
	"github.com/chizidotdev/nuntius/internal/core/service"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	config.LoadConfig()

	conn, err := gorm.Open(postgres.Open(config.EnvVars.PostgresUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	userStore := db.NewUserStore(conn)
	messageStore := db.NewMessageStore(conn)

	userService := service.NewUserService(userStore)
	messageService := service.NewMessageService(messageStore)

	server := drivers.NewController(
		userService,
		messageService,
	)

	port := config.EnvVars.PORT
	if port == "" {
		port = "8080"
	}

	serverAddr := "0.0.0.0:" + port
	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
