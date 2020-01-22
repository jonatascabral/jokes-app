package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Joke struct {
	gorm.Model
	Likes   int    `json:"likes" default:"0"`
	Unlikes int    `json:"unlikes" default:"0"`
	Joke    string `json:"joke" binding:"required" gorm:"type:varchar(100)"`
}

func (z Joke) New(joke string, likes int, unlikes int) *Joke {
	return &Joke{
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