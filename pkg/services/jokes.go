package services

import (
	"github.com/jonatascabral/jokes-app/pkg/models"
)

func GetJokes() (*[]*models.Joke, error) {
	return jokesRepository().GetJokes()
}

func GetJokeByID(jokeID int) (*models.Joke, error) {
	return jokesRepository().GetByID(jokeID)
}

func UpdateJoke(joke *models.Joke) (*models.Joke, error) {
	return jokesRepository().Save(joke)
}

func CreateJoke(joke *models.Joke) (*models.Joke, error) {
	return jokesRepository().Create(joke)
}