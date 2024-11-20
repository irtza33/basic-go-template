package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"errors"
	"github.com/irtza33/basic-go-template/config/DB"
	"github.com/irtza33/basic-go-template/infrastructure/database"
	"github.com/irtza33/basic-go-template/infrastructure/router"
)

func main() {
	fmt.Println("Hello, World!")

	dbConfig, err := DB.LoadConfig()
	if err != nil {
		panic(err)
	}

	sqlClient, err := database.NewSqlClient(context.Background(), dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Db)
	if err != nil {
		panic(err)
	}
	defer sqlClient.Conn.Close(context.Background())

	fmt.Println("Database successfully connected")

	router := router.NewGinRouter()

	go func() {
		fmt.Println("Server is now listening...")
		if err := router.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := router.Shutdown(); err != nil {
		panic(err.Error())
	}

}
