package models

type Joke struct {
	ID      int    `json:"id" binding:"required"`
	Likes   int    `json:"likes"`
	Unlikes int    `json:"unlikes"`
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
