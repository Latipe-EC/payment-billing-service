package message

import (
	"encoding/json"
	"latipe-payment-billing-service/app/config"
	order "latipe-payment-billing-service/app/data/dto"
	"latipe-payment-billing-service/app/service"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerOrderMessage struct {
	config    *config.Config
	pmService *service.PaymentService
}

func NewConsumerOrderMessage(config *config.Config, pmService *service.PaymentService) ConsumerOrderMessage {
	return ConsumerOrderMessage{
		config:    config,
		pmService: pmService,
	}
}

func (mq ConsumerOrderMessage) ListenOrderEventQueue(wg *sync.WaitGroup) {
	conn, err := amqp.Dial(mq.config.RabbitMQ.Connection)
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Printf("[%s] Comsumer has been connected", "INFO")

	channel, err := conn.Channel()
	defer channel.Close()
	defer conn.Close()

	// Khai báo một Exchange loại "direct"
	err = channel.ExchangeDeclare(
		mq.config.RabbitMQ.OrderEvent.Exchange, // Tên Exchange
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("cannot declare exchange: %v", err)
	}

	// Tạo hàng đợi
	_, err = channel.QueueDeclare(
		mq.config.RabbitMQ.OrderEvent.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("cannot declare queue: %v", err)
	}

	err = channel.QueueBind(
		mq.config.RabbitMQ.OrderEvent.Queue,
		mq.config.RabbitMQ.OrderEvent.RoutingKey,
		mq.config.RabbitMQ.OrderEvent.Exchange,
		false,
		nil)
	if err != nil {
		log.Fatalf("cannot bind exchange: %v", err)
	}

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		mq.config.RabbitMQ.OrderEvent.Queue, // queue
		mq.config.RabbitMQ.ConsumerName,     // consumer
		true,                                // auto ack
		false,                               // exclusive
		false,                               // no local
		false,                               // no wait
		nil,                                 //args
	)
	if err != nil {
		panic(err)
	}
	log.Printf("[queue:%s] waiting for messages...", mq.config.RabbitMQ.OrderEvent.Queue)

	// handle consumed messages from queue
	defer wg.Done()
	for msg := range msgs {
		log.Printf("[%s] received order message from: %s", "INFO", msg.RoutingKey)

		if err := mq.orderHandler(msg); err != nil {
			log.Printf("[%s] The order creation failed cause %s", "ERROR", err)
		}
	}

}

func (mq ConsumerOrderMessage) orderHandler(msg amqp.Delivery) error {
	message := order.OrderMessage{}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Printf("[%s] Parse message to order failed cause: %s", "ERROR", err)
		return err
	}

	if err := mq.pmService.CreatePaymentOfOrder(&message); err != nil {
		return err
	}
	
	return nil
}
