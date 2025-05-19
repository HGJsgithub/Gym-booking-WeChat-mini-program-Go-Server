package overdueOrder

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
)

func DepositOrderIntoMQ(orderID int64, conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return err
	}
	// 声明延迟队列（TTL + 死信交换机）
	args := amqp.Table{
		"x-message-ttl":             int32(0.5 * 60 * 1000), // TTL 1分钟（毫秒）
		"x-dead-letter-exchange":    "dlx_exchange",
		"x-dead-letter-routing-key": "order_cancel",
	}
	_, err = ch.QueueDeclare(
		"order_delay_queue", // 队列名
		true,                // 持久化
		false,               // 自动删除
		false,               // 排他
		false,               // 不等待
		args,                // 参数
	)
	if err != nil {
		return err
	}
	// 发布消息到延迟队列
	err = ch.Publish(
		"",                  // 使用默认交换机
		"order_delay_queue", // 路由键
		false,               // 强制
		false,               // 立即
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(strconv.FormatInt(orderID, 10)),
		},
	)
	return err
}
