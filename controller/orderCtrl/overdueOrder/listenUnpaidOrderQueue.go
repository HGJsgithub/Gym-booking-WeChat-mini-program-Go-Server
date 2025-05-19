package overdueOrder

import (
	"Gym_booking_WeChat_mini_program/controller/venueCtrl/changeVenueState"
	"Gym_booking_WeChat_mini_program/model"
	"context"
	"github.com/jinzhu/gorm"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"log"
)

func ListenUnpaidOrderQueue(coon *amqp.Connection, db *gorm.DB, redis *redis.Client) {
	ch, err := coon.Channel()
	if err != nil {
		log.Println("创建RabbitMQ Channel出错！", err)
		return
	}
	defer func(ch *amqp.Channel) {
		_ = ch.Close()
	}(ch)

	// 声明死信交换机和队列
	_ = ch.ExchangeDeclare(
		"dlx_exchange", // 交换机名
		"direct",       // 类型
		true,           // 持久化
		false,          // 自动删除
		false,          // 内部
		false,          // 不等待
		nil,
	)

	_, _ = ch.QueueDeclare(
		"order_cancel_queue", // 队列名
		true,                 // 持久化
		false,                // 自动删除
		false,                // 排他
		false,                // 不等待
		nil,
	)

	// 绑定队列到交换机
	_ = ch.QueueBind(
		"order_cancel_queue", // 队列名
		"order_cancel",       // 路由键
		"dlx_exchange",       // 交换机
		false,
		nil,
	)

	// 消费消息
	msgs, err := ch.Consume(
		"order_cancel_queue",
		"",
		false, // 关闭自动确认
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	for msg := range msgs {
		orderID := string(msg.Body)
		unpaidOrder := model.Order{}

		// 查询订单状态
		db.Table("orders").Where("id = ?", orderID).First(&unpaidOrder)
		if unpaidOrder.State == "待支付" {
			var updateReq changeVenueState.UpdateRequest
			//从redis获取要修改场地状态的信息
			err = redis.Get(ctx, orderID).Scan(&updateReq)
			if err != nil {
				log.Println("从redis获取场地状态更改条件失败！", err)
				log.Fatal(err)
			}
			updateInfo := updateReq.UpdateInfo
			//事务修改
			err = changeVenueState.Transaction(db, updateInfo, false)
			if err != nil {
				log.Println("更改场地状态失败！", err)
				log.Fatal(err)
			}
			unpaidOrder.State = "已取消"
			//将订单改为已取消
			db.Model(&unpaidOrder).Where("id = ?", orderID).Update("state", "已取消")
		}
		//删除对应订单的场地状态修改信息
		redis.Del(ctx, orderID)
		_ = msg.Ack(false)
	}
}
