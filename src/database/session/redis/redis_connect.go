package Session

import (
	Config "Riffy-music/src/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {

	Config.Init()

	fmt.Println("Подключение к Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     Config.RedisAddr(),
		Password: "",
		DB:       0,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
		return nil
	}
	fmt.Println("Связь с Redis установлена")

	return client

}
