package config

type Config struct {
	Postgres *Postgres
	Kafka    *Kafka
}

type Postgres struct {
	Host     string
	Port     uint
	User     string
	Password string
	Database string
}

type Kafka struct{}

func Load() (*Config, error) {
	// TOOD: Implement config.Load.
	panic("unimplemented")
}
