package router
 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"loremipsumbytes/app/controllers"
)
type Router interface {
	InstallRouter(app *fiber.App)
}

type HttpRouter struct {
    
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	 
  
	web := app.Group("", cors.New(), csrf.New())
	web.Get("/", controllers.RenderIndex)
  
  web.Post("/signup", controllers.SignupSubmit)
  web.Get("/gens", controllers.RenderGenerators)
  web.Get("/contact", controllers.RenderContact)  

  
  
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
