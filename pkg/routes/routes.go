package routes

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/jokes-app/pkg/controllers"
)

func LoadRoutes() *gin.Engine {
	router := gin.Default()

	loadStaticRoutes(router)
	loadApiRoutes(router)

	return router
}

func loadStaticRoutes(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./resources/views", true)))
}

func loadApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", controllers.Root)
		jokes := api.Group("/jokes")
		{
			jokes.GET("/", controllers.GetJokes)
			jokes.GET("/:jokeID", controllers.GetJoke)

			jokes.POST("/", controllers.CreateJoke)

			jokes.PUT("/like/:jokeID", controllers.LikeJoke)
			jokes.PUT("/unlike/:jokeID", controllers.UnlikeJoke)
		}
	}
}