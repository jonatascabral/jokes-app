package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/jokes-app/pkg/models"
	"github.com/jonatascabral/jokes-app/pkg/services"
	"log"
	"net/http"
	"strconv"
)

func acceptJson(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.Header("Accept", "application/json")
}

func Root(c *gin.Context) {
	acceptJson(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetJokes(c *gin.Context) {
	acceptJson(c)
	jokes, err := services.GetJokes()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, jokes)
}

func LikeJoke(c *gin.Context) {
	acceptJson(c)
	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		joke, err := services.GetJokeByID(jokeID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		joke.Likes += 1
		joke, err = services.UpdateJoke(joke)
		if err != nil {
			log.Fatal(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, &joke)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UnlikeJoke(c *gin.Context) {
	acceptJson(c)
	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		joke, err := services.GetJokeByID(jokeID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		joke.Unlikes += 1
		joke, err = services.UpdateJoke(joke)
		if err != nil {
			log.Fatal(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, joke)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func GetJoke(c *gin.Context) {
	acceptJson(c)
	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		joke, err := services.GetJokeByID(jokeID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, joke)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func CreateJoke(c *gin.Context) {
	acceptJson(c)
	joke := &models.Joke{}
	err := c.BindJSON(joke)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = services.CreateJoke(joke)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, joke)
}