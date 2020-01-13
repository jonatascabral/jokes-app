package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/jokes-app/pkg/services"
	"net/http"
	"strconv"
)

func acceptJson(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.Header("Accept", "application/json")
}

func Root(c *gin.Context) {
	acceptJson(c)
	c.JSON(http.StatusOK, "pong")
}

func GetJokes(c *gin.Context) {
	jokes := services.GetJokes()
	acceptJson(c)

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
		if joke.Likes > 0 {
			joke.Likes -= 1
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