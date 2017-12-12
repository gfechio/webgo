package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})

	f := os.Args[1] // Either Node:0000 or Way:000

	val2, err := client.Get(f).Result()
	if err == redis.Nil {
		fmt.Println(f, " does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(f, val2)
	}

	//	// Output: key value
	//	// key2 does not exist

}
