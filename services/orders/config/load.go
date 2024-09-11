package config

import "os"

func Load() (*Config, error) {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaGroup := os.Getenv("KAFKA_GROUP")

	orderEndedTopic := os.Getenv("KAFKA_TOPIC_ORDER_ENDED")
	acceptOrderTopic := os.Getenv("KAFKA_TOPIC_ACCEPT_ORDER")
	cookingStartedTopic := os.Getenv("KAFKA_TOPIC_COOKING_STARTED")
	createOrderTopic := os.Getenv("KAFKA_TOPIC_CREATE_ORDER")
	deliveryStartedTopic := os.Getenv("KAFKA_TOPIC_DELIVERY_STARTED")
	deliveryCompletedTopic := os.Getenv("KAFKA_TOPIC_DELIVERY_COMPLETED")
	orderHasBeenPayedTopic := os.Getenv("KAFKA_TOPIC_ORDER_PAYED")
	takeOrderTopic := os.Getenv("KAFKA_TOPIC_TAKE_ORDER")
	cancelOrderTopic := os.Getenv("KAFKA_TOPIC_CANCEL_ORDER")

	securityKey := os.Getenv("SECURITY_KEY")

	return &Config{
		Db: Postgres{
			Host:     postgresHost,
			Port:     postgresPort,
			User:     postgresUser,
			Password: postgresPassword,
			Db:       postgresDb,
		},
		MessageQueue: Kafka{
			Hosts: []string{kafkaHost},
			Group: kafkaGroup,
			Topics: topics{
				AcceptOrder:       acceptOrderTopic,
				CookingStarted:    cookingStartedTopic,
				CreateOrder:       createOrderTopic,
				DeliveryStarted:   deliveryStartedTopic,
				DeliveryCompleted: deliveryCompletedTopic,
				OrderHasBeenPayed: orderHasBeenPayedTopic,
				OrderEnded:        orderEndedTopic,
				TakeOrder:         takeOrderTopic,
				CancelOrder:       cancelOrderTopic,
			},
		},
		SecurityKey: securityKey,
	}, nil
}
