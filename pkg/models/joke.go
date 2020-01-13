package models

type Joke struct {
	ID 		int 	`json:"id" binding:"required"`
	Likes 	int 	`json:"likes"`
	Joke 	string 	`json:"joke" binding:"required"`
}

func (z Joke) New(id int, likes int, joke string) *Joke {
	return &Joke{
		ID: id,
		Likes: likes,
		Joke: joke,
	}
}