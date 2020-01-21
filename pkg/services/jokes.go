package services

import (
	"errors"
	"fmt"
	"github.com/jonatascabral/jokes-app/pkg/models"
)

var jokes = []*models.Joke{
	models.Joke{}.New(1, "Did you hear about the restaurant on the moon? Great food, no atmosphere.", 0, 0),
	models.Joke{}.New(2, "What do you call a fake noodle? An Impasta.", 0, 0),
	models.Joke{}.New(3, "How many apples grow on a tree? All of them.", 0, 0),
	models.Joke{}.New(4, "Want to hear a joke about paper? Nevermind it's tearable.", 0, 0),
	models.Joke{}.New(5, "I just watched a program about beavers. It was the best dam program I've ever seen.", 0, 0),
	models.Joke{}.New(6, "Why did the coffee file a police report? It got mugged.", 0, 0),
	models.Joke{}.New(7, "How does a penguin build it's house? Igloos it together.", 0, 0),
}

func GetJokes() *[]*models.Joke {
	//TODO Implement database logic
	return &jokes
}

func GetJokeByID(jokeID int) (*models.Joke, error) {
	//TODO Implement database logic
	jokes := GetJokes()
	for _, joke := range *jokes {
		if joke.ID == jokeID {
			return joke, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Joke with ID %d not found", jokeID))
}