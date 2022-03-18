package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/happierall/l"
	"github.com/hawkkiller/wtc-account/internal/database"
	"github.com/hawkkiller/wtc-account/internal/env"
	grpcApi "github.com/hawkkiller/wtc-account/transport/grpcApi"
	httpApi "github.com/hawkkiller/wtc-account/transport/httpApi"
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
	httpServer := httpApi.NewServerHTTP()
	grpcServer := grpcApi.NewServerGRPC()

	go func() {
		err := httpServer.StartServerHTTP()
		if err != nil {
			l.Print(err)
			os.Exit(1)
		}
	}()

	go func() {
		err := grpcServer.StartServerGRPC()
		if err != nil {
			l.Print(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

}
