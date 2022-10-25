package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mentorDelivery "autfinal/internal/microservice/mentor/delivery"
	mentorRepo "autfinal/internal/microservice/mentor/repository"
	mentorUsecase "autfinal/internal/microservice/mentor/usecase"

	childDelivery "autfinal/internal/microservice/child/delivery"
	childRepo "autfinal/internal/microservice/child/repository"
	childUsecase "autfinal/internal/microservice/child/usecase"

	scheduleDelivery "autfinal/internal/microservice/schedule/delivery"
	scheduleRepo "autfinal/internal/microservice/schedule/repository"
	scheduleUsecase "autfinal/internal/microservice/schedule/usecase"

	scheduleBeforeAfterDelivery "autfinal/internal/microservice/schedule_before_after/delivery"
	scheduleBeforeAfterRepo "autfinal/internal/microservice/schedule_before_after/repository"
	scheduleBeforeAfterUsecase "autfinal/internal/microservice/schedule_before_after/usecase"

	cardDelivery "autfinal/internal/microservice/card/delivery"
	cardRepo "autfinal/internal/microservice/card/repository"
	cardUsecase "autfinal/internal/microservice/card/usecase"

	cardBeforeAfterDelivery "autfinal/internal/microservice/card_before_after/delivery"
	cardBeforeAfterRepo "autfinal/internal/microservice/card_before_after/repository"
	cardBeforeAfterUsecase "autfinal/internal/microservice/card_before_after/usecase"

	imageDelivery "autfinal/internal/microservice/stock_image/delivery"
	imageRepo "autfinal/internal/microservice/stock_image/repository"
	imageUsecase "autfinal/internal/microservice/stock_image/usecase"

	log "autfinal/pkg/logger"

	"autfinal/internal/middleware"
	"autfinal/internal/router"
	"autfinal/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const logMessage = "server:"

func main() {
	message := logMessage + "Main:"
	log.Init(logrus.DebugLevel)
	log.Info(fmt.Sprintf(message+"started, log level = %s", logrus.DebugLevel))

	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Config isn't found 1")
		os.Exit(1)
	}
	viper.SetConfigFile("../../.env")
	err = viper.MergeInConfig()
	if err != nil {
		log.Error("Config isn't found 2")
		os.Exit(1)
	}

	postgresDB, err := utils.InitPostgres()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	mws := middleware.NewMiddleware()

	mentorR := mentorRepo.NewMentorRepository(postgresDB)
	childR := childRepo.NewChildRepository(postgresDB)
	scheduleR := scheduleRepo.NewScheduleRepository(postgresDB)
	scheduleBAR := scheduleBeforeAfterRepo.NewScheduleRepository(postgresDB)
	cardR := cardRepo.NewCardRepository(postgresDB)
	cardBAR := cardBeforeAfterRepo.NewBeforeAfterRepository(postgresDB)
	imageR := imageRepo.NewImageRepository(postgresDB)

	mentorU := mentorUsecase.NewMentorUsecase(mentorR)
	childU := childUsecase.NewChildUsecase(childR)
	scheduleU := scheduleUsecase.NewScheduleUsecase(scheduleR)
	scheduleBAU := scheduleBeforeAfterUsecase.NewScheduleBeforeAfterUsecase(scheduleBAR)
	cardU := cardUsecase.NewCardUsecase(scheduleR, cardR)
	cardBAU := cardBeforeAfterUsecase.NewCardBeforeAfterUsecase(scheduleBAR, cardBAR)
	imageU := imageUsecase.NewImageUsecase(imageR)

	mentorD := mentorDelivery.NewMentorDelivery(mentorU)
	childD := childDelivery.NewChildDelivery(childU)
	scheduleD := scheduleDelivery.NewScheduleDelivery(scheduleU)
	scheduleBAD := scheduleBeforeAfterDelivery.NewScheduleBeforeAfterDelivery(scheduleBAU)
	cardD := cardDelivery.NewCardDelivery(cardU)
	cardBAD := cardBeforeAfterDelivery.NewCardBeforeAfterDelivery(cardBAU)
	imageD := imageDelivery.NewImageDelivery(imageU)


	baseRouter := gin.New()
	baseRouter.Use(gin.Logger())
	baseRouter.Use(gin.Recovery())
	baseRouter.Use(mws.CORSMiddleware())
	baseRouter.MaxMultipartMemory = 8 << 20

	routerAPI := baseRouter.Group("/api")

	mentorRouter := routerAPI.Group("/mentors")
	router.MentorEndpoints(mentorRouter,mentorD)

	childRouter := routerAPI.Group("/childs")
	router.ChildEndpoints(childRouter, childD)

	scheduleDayRouter := childRouter.Group("/:child_id/schedules/day")
	router.ScheduleDayEndpoints(scheduleDayRouter, scheduleD)

	scheduleLessonRouter := childRouter.Group("/:child_id/schedules/lesson")
	router.ScheduleLessonEndpoints(scheduleLessonRouter, scheduleD)

	scheduleBeforeAfterRouter := childRouter.Group("/:child_id/schedules/before_after")
	router.ScheduleBeforeAfterEndpoints(scheduleBeforeAfterRouter, scheduleBAD)

	cardDayRouter := routerAPI.Group("/schedules/day/:schedule_id/cards")
	router.CardDayEndpoints(cardDayRouter, cardD)

	cardLessonRouter := routerAPI.Group("/schedules/lesson/:schedule_id/cards")
	router.CardLessonEndpoints(cardLessonRouter, cardD)

	cardBeforeAfterRouter := routerAPI.Group("/schedules/before_after/:schedule_id/cards")
	router.CardBeforeAfterEndpoints(cardBeforeAfterRouter, cardBAD)

	imageRouter := routerAPI.Group("/stock")
	router.ImagesEndpoints(imageRouter, imageD)

	port := viper.GetString("server.port")

	server := &http.Server{
		Addr: ":"+port,
		Handler: baseRouter,
		IdleTimeout: 10 * time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <- sigChan
	log.Info("Graceful shutdown:", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(timeoutContext)

}