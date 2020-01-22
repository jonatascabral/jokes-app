package repositories

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jonatascabral/jokes-app/pkg/models"
)

type JokesRepository interface {
	Create(joke *models.Joke) (*models.Joke, error)
	Save(joke *models.Joke) (*models.Joke, error)
	GetByID(jokeID int) (*models.Joke, error)
	Delete(jokeID int) (bool, error)
	GetJokes() (*[]*models.Joke, error)
}

type jokesRepository struct {
	client *gorm.DB
}

func NewJokesRepository (client *gorm.DB) JokesRepository {
	client.AutoMigrate(&models.Joke{})
	return &jokesRepository{
		client: client,
	}
}

func (repository *jokesRepository) GetJokes() (*[]*models.Joke, error) {
	var jokes []*models.Joke
	err := repository.client.Find(&jokes).Error

	if err != nil {
		return &jokes, err
	}

	return &jokes, nil
}

func (repository *jokesRepository) GetByID(jokeID int) (*models.Joke, error) {
	var joke models.Joke
	err := repository.client.Find(&joke, jokeID).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Joke with ID %d not found", jokeID))
	}

	return &joke, nil
}

func (repository *jokesRepository) Save(joke *models.Joke) (*models.Joke, error) {
	err := repository.client.Save(joke).Error
	if err != nil {
		return nil, err
	}
	return joke, nil
}

func (repository *jokesRepository) Create(joke *models.Joke) (*models.Joke, error) {
	err := repository.client.Create(joke).Error
	if err != nil {
		return nil, err
	}
	return joke, nil
}

func (repository *jokesRepository) Delete(jokeID int) (bool, error) {
	joke, err := repository.GetByID(jokeID)
	if err != nil {
		return false, err
	}

	err = repository.client.Delete(joke).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
