package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"loremipsumbytes/pkg/hasura" 
	"net/http"
	"time" 
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2" 
)

type DataForm struct {
	Lor_Users      []User       `json:"lor_users"`
	Lor_Gens       []Generators `json:"lor_gens"`
}

type Generators struct {
	Title string `json:"title"`
	Link string `json:"link"`
	Content string `json:"content"`
}

type User struct {
	Full_name string `json:"full_name"`
	Email     string `json:"email"`
	Changed_on string `json:"changed_on"`
	Created_on string `json:"created_on"`
	Password string `json:"password"`
	Id string `json:"id"`
}

type Request struct {
	Data DataForm
}

var Url = "https://loremdb.hasura.app/v1/graphql"

var Store = session.New()

func RenderLogout(c *fiber.Ctx) error {
	
	sess, err := Store.Get(c)
	if err != nil {
		log.Println(err)
	}

	sess.Delete("name")
 
	// Destry session
	if err := sess.Destroy(); err != nil {
		log.Println("Session not destroyed error")
	}

	return c.Redirect("/")
 
}
func RenderHome(c *fiber.Ctx) error {
	return c.Redirect("/")
}
func RenderIndex(c *fiber.Ctx) error {
	sess, err :=Store.Get(c)
	log.Println("Index Sess--");log.Println(sess)
	if err != nil {
		log.Println("Index err")
		log.Println(err)
	}

	q := hasura.Query_user() 
 	data_resp := fetchHttp(q) 
 
	name := sess.Get("name");log.Println(sess.Get("name"))
	showLogin := name != nil
	log.Println(sess);log.Println(name);log.Printf("- %v" , showLogin); // this is bool
  
	gens := fetchGenerators()
	log.Println(data_resp.Lor_Users);log.Println("gens");log.Println("gens");log.Println(gens);

	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Lorem Generators",
		"Loremusers": data_resp.Lor_Users,
		"Loremgens": gens, 
		"Namelogin": name, 
	},  "layouts/template") 
}

func fetchGenerators() []Generators {

	q := hasura.Query_gens()
	data_resp := fetchHttp(q);log.Println("Fetch Gen Problem");log.Println("Fetch Gen Problem");
	log.Println(data_resp);log.Println(data_resp);
	return data_resp.Lor_Gens
}

func RenderAddGraph(c *fiber.Ctx) error {

	q := hasura.Query_user()
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(q))
	if err != nil {
		log.Println("Error in Gql Run")
	}

	client := &http.Client{
		Timeout: time.Second * 7,
	}
	response, err := client.Do(req)

	if err != nil {
		log.Println("Error in Clieee snt Do Run")
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	log.Println(string(data))

	data_resp := Request{}

	if err := json.Unmarshal(data, &data_resp); err != nil {
		log.Println("Can not unmnarshal JSON")
	}
	// log.Println("This -> HTTP Response: %s")
	// log.Println(data_resp.Data.Loremsite_Users[0])

	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/template")
}
func AddGens(c *fiber.Ctx) error {

	return c.Render("meme", fiber.Map{
		"FiberTitle": "Meme Gen",
	}, "layouts/generic")
}
func RenderMeme(c *fiber.Ctx) error {

	return c.Render("meme", fiber.Map{
		"FiberTitle": "Meme Gen",
	}, "layouts/generic")
}
func RenderFantasy(c *fiber.Ctx) error {

	return c.Render("fantasy", fiber.Map{
		"FiberTitle": "Fantasy Map",
	}, "layouts/generic")
}
func RenderLogin(c *fiber.Ctx) error {
    sess, err := Store.Get(c)
	if err != nil {
		log.Println(sess)
		log.Println("Error in Clieee snt Do Run")
	}
	return c.Render("login", fiber.Map{
		"SingleTitle": "Login",
	}, "layouts/generic")
}
func RenderRegister(c *fiber.Ctx) error {

	return c.Render("register", fiber.Map{
		"FiberTitle": "Fantasy Map",
	}, "layouts/generic")
}
   
func SignupSubmit(c *fiber.Ctx) error { 
	sess, err := Store.Get(c)
	full_name := c.FormValue("name")
	email := c.FormValue("email") 
	password := c.FormValue("password")

	created_on := "2022-01-01 00:00:00"; changed_on :="2022-01-01 00:00:00" ;

	q := hasura.Mutation_signup_user( password , full_name , email ,created_on  ,changed_on   ) 
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(q))
 

	if err != nil {
		log.Println("Error in  Mutation")
		log.Println(err)
	}else{
		sess.Set("name", full_name)
		name := sess.Get("name") 
	    log.Printf("Login name %v" , name)
		
		return c.Redirect("/home")
 
	} 
	 
	log.Println("Req in ---SignupSubmit Mutation" , req)

	log.Println("Failed to sign up")
	c.Redirect("/gens")
	return c.Redirect("/")

}
func ReadGen(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		log.Println("Error in Clieee snt Do Run")
		log.Println(sess)

	}
	n := hasura.Query_gens_byname()
	log.Println(n)

	return c.Render("readgenerator", fiber.Map{
		"n": "Hello From Fiber Html Engine",
	}, "layouts/template")
}

func RenderGenerators(c *fiber.Ctx) error {

	gens := fetchGenerators()

	return c.Render("listgens", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
		"Loremgens":  gens,
	}, "layouts/template")
}

func RenderContact(c *fiber.Ctx) error {

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/template")
}

func fetchHttp(hq []byte) DataForm {
	// Hasura query in bytes  hq [] byte
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(hq))
	if err != nil {
		log.Println("Error in Gql Run")
	}

	client := &http.Client{
		Timeout: time.Second * 7,
	}
	response, err := client.Do(req)

	defer response.Body.Close()
	if err != nil {
		log.Println("Error in Clieee snt Do Run")
	}
	data_bytes, _ := ioutil.ReadAll(response.Body)
	log.Println("string(data_bytes)");log.Println("string(data_bytes)");
	log.Println(string(data_bytes))
	data_resp := Request{}
	err = json.Unmarshal(data_bytes, &data_resp)
	if err != nil {
		log.Println("Can n not unmnarshal JSON")
	}
	return data_resp.Data
}
