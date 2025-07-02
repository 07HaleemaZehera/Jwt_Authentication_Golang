package routes

import (
	"JWT_AUTHENTICATION/controllers"

	"github.com/gofiber/fiber/v2"
)
func SetUpRoutes(app *fiber.App){
app.Get("/", controllers.Hello)
}