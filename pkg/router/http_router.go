package router
 
import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"loremipsumbytes/app/controllers"
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

  
  
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
