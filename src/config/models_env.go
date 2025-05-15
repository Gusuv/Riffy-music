package Config

import (
	"os"
)

func RedisAddr() string {
	var Redis_addr = os.Getenv("REDIS_ADDR")
	return Redis_addr
}

func DatabaseDns() string {
	db_dns := os.Getenv("DATABASE_DSN")
	return db_dns
}
