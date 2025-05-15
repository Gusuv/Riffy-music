package Session

import (
	Config "Riffy-music/src/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConnectRedis() {

	Config.Init()

	fmt.Println("Подключение к Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     Config.RedisAddr(),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Подключение удалось:", pong)

	defer client.Close()

}
