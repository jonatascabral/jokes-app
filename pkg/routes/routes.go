package routes

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/jokes-app/pkg/controllers"
)

func LoadRoutes(router *gin.Engine) {
	loadStaticRoutes(router)
	loadApiRoutes(router)
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

			jokes.POST("/like/:jokeID", controllers.LikeJoke)
			jokes.POST("/unlike/:jokeID", controllers.UnlikeJoke)
		}
	}
}