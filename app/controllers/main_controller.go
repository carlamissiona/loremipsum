package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"loremipsumbytes/pkg/hasura"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DataForm struct {
	Loremsite_Users      []User       `json:"loremsite_users"`
	Loremsite_Generators []Generators `json:"loremsite_generators"`
}

type Generators struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Desc string `json:"desc"`
}

type User struct {
	Full_name string `json:"full_name"`
	Email     string `json:"email"`
}

type Request struct {
	Data DataForm
}

var Url = "https://loremdb.hasura.app/v1/graphql"

func RenderIndex(c *fiber.Ctx) error {
	q := hasura.Query_user()
	data_resp := fetchHttp(q)

	gens := fetchGenerators()
	log.Println(data_resp.Loremsite_Users[0])
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
		"Loremusers": data_resp.Loremsite_Users,
		"Loremgens":  gens,
	}, "layouts/htm")
}

func fetchGenerators() []Generators {

	q := hasura.Query_gens()
	data_resp := fetchHttp(q)
	log.Println(data_resp.Loremsite_Generators[0])

	return data_resp.Loremsite_Generators
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
	}, "layouts/htm")
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

	email := c.Params("email")
	pass := c.Params("password")

	q := hasura.Mutation_signup_user(email, pass)
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(q))
	if err != nil {
		log.Println("Error in  Mutation")
	}
	log.Println("Success in  Mutation")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Println("Error in Resp Body Client Run")
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	log.Println(string(data))
	c.Redirect("/")
	return nil

}

func FormGens(c *fiber.Ctx) error {

	return c.Render("readgenerator", fiber.Map{
		"n": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func AddGens(c *fiber.Ctx) error {

	desc := c.Params("description_ofgenerator")
	name := c.Params("name_ofgenerator")
	link := c.Params("link")
	g := hasura.Mutation_add_gen(desc, link, name)
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(g))
	if err != nil {
		log.Println("Error in Add Generator Mutation")
	}
	log.Println("Success in  Mutation Generator")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error in Resp Body Client Run")
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	log.Println(string(data))
	c.Redirect("/gens")
	return nil

}
func ReadGen(c *fiber.Ctx) error {
	n := hasura.Query_gens_byname()
	log.Println(n)

	return c.Render("readgenerator", fiber.Map{
		"n": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderGenerators(c *fiber.Ctx) error {

	gens := fetchGenerators()

	return c.Render("listgens", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
		"Loremgens":  gens,
	}, "layouts/htm")
}

func RenderContact(c *fiber.Ctx) error {

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
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

	log.Println(string(data_bytes))
	data_resp := Request{}
	err = json.Unmarshal(data_bytes, &data_resp)
	if err != nil {
		log.Println("Can n not unmnarshal JSON")
	}
	return data_resp.Data
}
