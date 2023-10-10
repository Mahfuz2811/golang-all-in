package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mahfuz/golang-all-in/packages/db"

	"github.com/mahfuz/golang-all-in/packages/redis"
)

func main() {

	db, err := db.DbConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()

	log.Printf("Successfully connected to database")

	redisClient, err := redis.RedisConnection()
	if err != nil {
		log.Printf("Failed to connect to redis: %s", err)
		return
	}

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)

	err = godotenv.Load()
	if err != nil {
		log.Printf("Failed to load env: %s", err)
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println(s3Bucket, secretKey)

	// value, ok := config.Config("DB_USER", "User")
	// fmt.Println(value, ok)
}
