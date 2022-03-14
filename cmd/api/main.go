package main

import (
	"fmt"
	"github.com/happierall/l"
	"github.com/hawkkiller/wtc-account/api"
	"github.com/hawkkiller/wtc-account/internal/database"
	"github.com/hawkkiller/wtc-account/internal/env"
	"log"
	"os"
	"os/signal"
)

// @title WTC ACCOUNT SERVICE
// @version 0.8+07-01-2022-22:30
// @description Account service for WTC.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email miskadl09@gmail.com

// @host localhost:9000
// @BasePath /api/v1/account-service
// @schemes http
func main() {
	env.SetupEnv()
	database.SetupDB()
	server := api.SetupApi()

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "9000"
		}
		err := server.Start(fmt.Sprintf(":%s", port))
		if err != nil {
			l.Print(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

}