package main

import (
	"fmt"
	"latipe-payment-billing-service/app/config"
	"latipe-payment-billing-service/app/message"
	"latipe-payment-billing-service/app/repository"
	"latipe-payment-billing-service/app/repository/gorm"
	"latipe-payment-billing-service/app/service"
	"sync"
)

func main() {
	fmt.Println("Init billing & payment worker")

	//read config file
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	//database connection
	dbCnt := gorm.NewMySQLConnection(cfg)
	//repository
	err = repository.InitCommissionLevelData(dbCnt)
	if err != nil {
		panic(err.Error())
	}

	paymentRepo := repository.NewPaymentRepository(dbCnt)
	billingRepo := repository.NewBillingRepository(dbCnt)

	//service
	paymentService := service.NewPaymentService(&paymentRepo)
	billingSevice := service.NewBillingService(&paymentRepo, &billingRepo)

	//consumer
	orderConsumer := message.NewConsumerOrderMessage(cfg, &paymentService)
	storeConsumer := message.NewConsumerStoreMessage(cfg, &paymentService, &billingSevice)

	//init sync goroutine group
	var wg sync.WaitGroup

	wg.Add(1)
	go orderConsumer.ListenOrderEventQueue(&wg)

	wg.Add(1)
	go storeConsumer.ListenOrderEventQueue(&wg)

	wg.Wait()
}
