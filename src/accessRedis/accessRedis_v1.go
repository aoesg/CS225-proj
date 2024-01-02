package accessRedis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Set_v1(redis_address string, key string, value string) error {

	client := redis.NewClient(&redis.Options{
		Addr:     redis_address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(client)

	ctx := context.Background()

	err := client.Set(ctx, key, value, 0).Err()
	checkError(err)

	fmt.Printf("GET %s : %s\n", key, value)

	return err
}

func Get_v1(redis_address string, key string) (string, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     redis_address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(client)

	ctx := context.Background()

	value, err := client.Get(ctx, key).Result()
	checkError(err)

	fmt.Printf("SET %s : %s\n", key, value)

	return value, err
}
