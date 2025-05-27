package messageQueue

import (
	"Gym_booking_WeChat_mini_program/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ() *amqp.Connection {
	var conf config.Config
	conf.LoadConfig()
	rabbitMQ := conf.RabbitMQ
	args := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQ.Admin, rabbitMQ.Password, rabbitMQ.Host, rabbitMQ.Port)
	conn, err := amqp.Dial(args)
	if err != nil {
		fmt.Println("连接RabbitMQ出错！", err)
		return nil
	}
	return conn
}
