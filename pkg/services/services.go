package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jonatascabral/jokes-app/pkg/repositories"
	"log"
	"os"
	"strconv"
	"time"
)

func jokesRepository() (repositories.JokesRepository) {
	client := ConnectDatabase()
	return repositories.NewJokesRepository(client)
}

func ConnectCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		ReadTimeout: time.Duration(10 * time.Second),
	})

	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	return client
}

func ConnectDatabase() *gorm.DB {
	var connectionString string
	driver := os.Getenv("DATABASE_DRIVER")

	switch driver {
	case "mysql":
		connectionString = fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_NAME"),
			"utf8",
		)
	case "postgres":
		connectionString = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_PASSWORD"),
		)
	}
	database, err := gorm.Open(driver, connectionString)
	if err != nil {
		panic(err)
	}
	logMode, _ := strconv.ParseBool(os.Getenv("DATABASE_LOG"))
	log.Println(logMode)
	database.LogMode(logMode)

	return database
}
