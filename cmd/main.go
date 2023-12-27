package main

import (
	"github.com/chizidotdev/nuntius/config"
	"github.com/chizidotdev/nuntius/internal/app/db"
	"github.com/chizidotdev/nuntius/internal/core/service"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	conn, err := gorm.Open(postgres.Open(config.EnvVars.PostgresUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	userStore := db.NewUserStore(conn)
	messageStore := db.NewMessageStore(conn)

	userService := service.NewUserService(userStore)
	messageService := service.NewMessageService(messageStore)

	server := http.NewHTTPServer(
		userService,
		messageService,
	)

	port := config.EnvVars.PORT
	if port == "" {
		port = "8080"
	}

}
