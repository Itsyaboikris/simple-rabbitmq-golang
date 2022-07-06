package main

import (
	"fmt"
	"rabbitmq-crash-course/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go RabbitMQ Crash Course")

	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error Setting Up Application")
		fmt.Println(err)
	}
}
