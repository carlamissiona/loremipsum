package router

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"loremipsumbytes/app/controllers"

	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type Router interface {
	InstallRouter(app *fiber.App)
}

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {

	app.Use(csrf.New())

	app.Get("/", controllers.RenderIndex)

	app.All("/signup", controllers.SignupSubmit)
	app.Get("/gens", controllers.RenderGenerators)
	app.Get("/contact", controllers.RenderContact)
	app.Get("/gens-forms", controllers.AddGens)

}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
