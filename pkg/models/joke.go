package models

import "encoding/json"

type Joke struct {
	ID      int    `json:"id"`
	Likes   int    `json:"likes" default:"0"`
	Unlikes int    `json:"unlikes" default:"0"`
	Joke    string `json:"joke" binding:"required"`
}

func (z Joke) New(id int, joke string, likes int, unlikes int) *Joke {
	return &Joke{
		ID:      id,
		Likes:   likes,
		Unlikes: unlikes,
		Joke:    joke,
	}
}

func (z Joke) FromJSON(jokeJson string) (*Joke, error) {
	joke := &Joke{}
	err := json.Unmarshal([]byte(jokeJson), joke)
	return joke, err
}

func (z Joke) ToJSON() (string, error) {
	jokeJson, err := json.Marshal(z)
	return string(jokeJson), err
}