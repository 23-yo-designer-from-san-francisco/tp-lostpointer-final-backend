package main

import (
	"log"
	"os"
	"time"
	"os/signal"
	"syscall"
	"net/http"
	"context"

	cardDelivery "autfinal/internal/microservice/card/delivery"
	cardRepo "autfinal/internal/microservice/card/repository"
	cardUsecase "autfinal/internal/microservice/card/usecase"

	"autfinal/internal/router"
	"autfinal/internal/middleware"
	"autfinal/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	l := log.New(os.Stdout, "Diploma-API", log.LstdFlags)

	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Config isn't found 1")
		os.Exit(1)
	}
	viper.SetConfigFile("../../.env")
	err = viper.MergeInConfig()
	if err != nil {
		log.Print("Config isn't found 2")
		os.Exit(1)
	}

	postgresDB, err := utils.InitPostgres()
	if err != nil {
		log.Print("InitPG")
		log.Println(err)
		os.Exit(1)
	}

	mws := middleware.NewMiddleware()

	cardR := cardRepo.NewCardRepository(postgresDB)

	cardU := cardUsecase.NewCardUsecase(cardR)

	cardD := cardDelivery.NewCardDelivery(cardU)

	baseRouter := gin.New()
	baseRouter.Use(gin.Logger())
	baseRouter.Use(gin.Recovery())
	baseRouter.MaxMultipartMemory = 8 << 20

	routerAPI := baseRouter.Group("/api")

	userRouter := routerAPI.Group("/cards")
	router.CardEndpoints(userRouter, mws, cardD)

	port := viper.GetString("server.port")

	server := &http.Server{
		Addr: ":"+port,
		ErrorLog: l,
		Handler: baseRouter,
		IdleTimeout: 10 * time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <- sigChan
	log.Println("Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(timeoutContext)

}