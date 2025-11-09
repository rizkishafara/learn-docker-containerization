package routes

import (
	"src/handlers"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	app.Get("/", handlers.DashboardPage)
}
func Users(app *fiber.App) {
	app.Get("/users", handlers.UsersPage)
	app.Post("/users", handlers.UserCreate)
	app.Post("/users/update/:id", handlers.UserUpdate)
	app.Get("/users/delete/:id", handlers.UserDelete)
}
func Products(app *fiber.App) {
	app.Get("/products", handlers.ProductsPage)
	app.Post("/products", handlers.ProductCreate)
	app.Post("/products/update/:id", handlers.ProductUpdate)
	app.Get("/products/delete/:id", handlers.ProductDelete)
}
