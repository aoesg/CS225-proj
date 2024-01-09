package accessRedis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func Set_v1(redis_address string, key string, value string) error {

	client := redis.NewClient(&redis.Options{
		Addr: redis_address,
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})
	// fmt.Println(client)

	ctx := context.Background()

	err := client.Set(ctx, key, value, 0).Err()
	checkError(err)

	fmt.Printf("GET %s : %s\n", key, value)

	return err
}

func Get_v1(redis_address string, key string) (string, error) {

	client := redis.NewClient(&redis.Options{
		Addr: redis_address,
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})
	// fmt.Println(client)

	ctx := context.Background()

	value, err := client.Get(ctx, key).Result()
	checkError(err)

	fmt.Printf("SET %s : %s\n", key, value)

	return value, err
}

func Incr_v1(redis_address string, key string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redis_address,
	})
	fmt.Println(client)

	ctx := context.Background()

	value_interface, err := client.Do(ctx, "INCR", key).Result()
	checkError(err)

	value_int64 := value_interface.(int64)
	value := strconv.FormatInt(value_int64, 10)

	fmt.Printf("INCR %s to %s\n", key, value)

	return value, err
}

func DelPattern(redis_address string, pattern string) (int64, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redis_address,
	})
	fmt.Println(client)

	ctx := context.Background()

	keys, err := client.Keys(ctx, pattern).Result()
	checkError(err)

	if len(keys) == 0 {
		return 0, err
	}

	// keys_inOneStr := strings.Join(keys, " ")

	delCount, err := client.Del(ctx, keys[:]...).Result() // 新知识新知识
	checkError(err)

	return delCount, err
}
