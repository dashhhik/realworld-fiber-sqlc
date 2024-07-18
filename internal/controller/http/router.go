package http

import (
	"github.com/gofiber/fiber/v2"
	"realworld-fiber-sqlc/internal/controller/http/handlers"
	"realworld-fiber-sqlc/pkg/middleware"
	"realworld-fiber-sqlc/usecase/dto/sqlc"
)

func SetupRoutes(app *fiber.App, dbQueries *sqlc.Queries) {

	handlerBase := handlers.NewHandlerQ(dbQueries)

	api := app.Group("/api")

	users := api.Group("/users")

	users.Post("/login", handlerBase.Login)
	users.Post("/", handlerBase.Register)

	user := api.Group("/user")

	user.Get("/", middleware.Protected(), handlerBase.CurrentUser)
	user.Put("/", middleware.Protected(), handlerBase.UpdateProfile)
	//

	profilesRoute := api.Group("/profiles")
	profilesRoute.Get("/:username", handlerBase.GetProfile)
	//
	//articlesRoute.Get("/")

	articlesRoute := api.Group("/articles")

	articlesRoute.Get("/:slug", handlerBase.GetArticle)
	//
	app.Get("api/tags", handlerBase.GetTags)

	//user

	//profilesRoute
	//
	profilesRoute.Post("/:username/follow", middleware.Protected(), handlerBase.Follow)
	profilesRoute.Delete("/:username/follow", middleware.Protected(), handlerBase.Unfollow)
	//
	////	articles

	//articlesRoute.Get("/feed")
	//
	articlesRoute.Post("/", middleware.Protected(), handlerBase.CreateArticle)
	articlesRoute.Put("/:slug", middleware.Protected(), handlerBase.UpdateArticle)
	articlesRoute.Delete("/:slug", middleware.Protected(), handlerBase.DeleteArticle)
	//
	//comments
	//
	commentsRoute := articlesRoute.Group("/:slug/comments")
	//commentsRoute := articlesRoute.Group("/:slug/comments")
	commentsRoute.Post("/", middleware.Protected(), handlerBase.CreateComment)
	//commentsRoute.Get("/")
	//commentsRoute.Delete("/:id")
	//
	//ffv
	app.Post("/api/articles/:slug/favorite", middleware.Protected(), handlerBase.FavoriteArticle)
	app.Delete("/api/articles/:slug/favorite", middleware.Protected(), handlerBase.UnfavoriteArticle)

}
