package router

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"loremipsumbytes/app/controllers"
 
	// "log"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
) 

type Router interface {
	InstallRouter(app *fiber.App)
}

type HttpRouter struct {
}
 
func (h HttpRouter) InstallRouter(app *fiber.App) {
    
    
	app.Get("/", controllers.RenderIndex)
	app.Get("/logout", controllers.RenderLogout)
	app.All("/signup", controllers.SignupSubmit) 
	app.All("/home", controllers.RenderHome) 
	app.Get("/login", controllers.RenderLogin)
	app.Get("/register", controllers.RenderRegister)
	app.Get("/gens", controllers.RenderGenerators)
	app.Get("/contact", controllers.RenderContact)
	app.Get("/gens-forms", controllers.AddGens)
	app.Get("/static/fantasymap", controllers.RenderFantasy)
	 
 
	 
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
