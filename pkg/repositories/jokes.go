package repositories

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jonatascabral/jokes-app/pkg/models"
	"log"
	"sort"
	"strconv"
	"strings"
)

var (
	keyAllJokes     = "jokes-*"
	keyJokes = "jokes-"
	keyJokesPattern = keyJokes + "%d"
)

type JokesRepository interface {
	Create(joke *models.Joke) (*models.Joke, error)
	Save(joke *models.Joke) (*models.Joke, error)
	GetByID(jokeID int) (*models.Joke, error)
	Delete(jokeID int) (bool, error)
	GetJokes() (*[]*models.Joke, error)
}

type jokesRepository struct {
	client *redis.Client
}

func NewJokesRepository (client *redis.Client) JokesRepository {
	return &jokesRepository{
		client: client,
	}
}

func (repository *jokesRepository) GetJokes() (*[]*models.Joke, error) {
	jokesIDs, err := repository.client.Keys(keyAllJokes).Result()

	if err != nil {
		return &[]*models.Joke{}, err
	}

	jokes := []*models.Joke{}
	for _, jokeID := range jokesIDs {
		jokeJson, _ := repository.client.Get(jokeID).Result()

		joke, err := models.Joke.FromJSON(models.Joke{}, jokeJson)
		if err != nil {
			return &[]*models.Joke{}, err
		}

		jokes = append(jokes, joke)
	}
	return &jokes, nil
}

func (repository *jokesRepository) GetByID(jokeID int) (*models.Joke, error) {
	jokeJson, err := repository.client.Get(fmt.Sprintf(keyJokesPattern, jokeID)).Result()
	if err != nil {
		return nil, err
	}
	if jokeJson != "" {
		return models.Joke.FromJSON(models.Joke{}, jokeJson)
	}
	return nil, errors.New(fmt.Sprintf("Joke with ID %d not found", jokeID))
}

func (repository *jokesRepository) Save(joke *models.Joke) (*models.Joke, error) {
	json, err := joke.ToJSON()
	if err != nil {
		return joke, err
	}
	_, err = repository.client.Set(fmt.Sprintf(keyJokesPattern, joke.ID), json, 0).Result()
	if err != nil {
		return nil, err
	}
	return joke, nil
}

func (repository *jokesRepository) Create(joke *models.Joke) (*models.Joke, error) {
	jokesIDs, err := repository.client.Keys(keyAllJokes).Result()
	jokesIDsSlice := sort.StringSlice(jokesIDs)
	sort.Sort(sort.Reverse(jokesIDsSlice))
	log.Println(jokesIDsSlice)
	if err != nil {
		return nil, err
	}

	lastJokeID, err := strconv.Atoi(strings.Split(jokesIDsSlice[0], keyJokes)[1])
	if err != nil {
		return nil, err
	}

	joke.ID = lastJokeID + 1
	json, err := joke.ToJSON()
	if err != nil {
		return nil, err
	}

	_, err = repository.client.Set(fmt.Sprintf(keyJokesPattern, joke.ID), json, 0).Result()
	if err != nil {
		return nil, err
	}
	return joke, nil
}

func (repository *jokesRepository) Delete(jokeID int) (bool, error) {
	_, err := repository.client.Del(fmt.Sprintf(keyJokesPattern, jokeID)).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
