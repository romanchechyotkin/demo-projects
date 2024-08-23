package main

import (
	"log"
	redisapi "mistakes/14_package_name_collisions/redis"
)

func main() {
	redis := redisapi.NewClient()
	redis.Log()
	// to avoid name collision use alias for import

	// avoid naming collisions between a variable and a built-in keywords
	copy := copyFile()
	log.Println(copy)
}

func copyFile() string {
	log.Println("copy")
	return ""
}
