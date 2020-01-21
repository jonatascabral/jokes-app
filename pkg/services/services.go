package services

import (
	"github.com/jonatascabral/jokes-app/pkg/repositories"
)

func jokesRepository() (repositories.JokesRepository) {
	client := ConnectRedis()
	return repositories.NewJokesRepository(client)
}