package redis

type Config struct {
	Addr string `env:"REDIS_ADDR"`
}

var defaultConfig = Config{
	Addr: "localhost:6379",
}
