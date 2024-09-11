package config

type (
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		Db       string
	}

	topics struct {
		AcceptOrder       string
		CookingStarted    string
		CreateOrder       string
		DeliveryStarted   string
		DeliveryCompleted string
		OrderHasBeenPayed string
		OrderEnded        string
		TakeOrder         string
		CancelOrder       string
	}

	Kafka struct {
		Hosts []string
		Group string

		Topics topics
	}

	Config struct {
		Db           Postgres
		MessageQueue Kafka

		SecurityKey string
	}
)
