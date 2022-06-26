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
	//*********/ Query User
	q := hasura.Query_user()
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(q))
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
	log.Println("This ---> HTTP Response: %s")
	log.Println(data_resp.Data.Loremsite_Users[0])

	gens := fetchGenerators()

	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
		"Loremusers": data_resp.Data.Loremsite_Users,
		"Loremgens":  gens,
	}, "layouts/htm")
}

func fetchGenerators() []Generators {

	q := hasura.Query_gens()
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(q))
	if err != nil {
		log.Println("Error in Gql Run")
	}

	client := &http.Client{
		Timeout: time.Second * 7,
	}
	response, err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		log.Println("Error in Resp Body  Do Run")
	}
	data_bytes, _ := ioutil.ReadAll(response.Body)

	log.Println(string(data_bytes))

	data_resp := Request{}
	err = json.Unmarshal(data_bytes, &data_resp)

	if err != nil {
		log.Println("Can not unmnarshal JSON")
	}
	log.Println("This -> HTTP Response: Loremsite_Generators %s")
	log.Println(data_resp.Data.Loremsite_Generators[0])

	return data_resp.Data.Loremsite_Generators
}

func RenderAddGraph(c *fiber.Ctx) error {
	//**********/ Query User
	q := hasura.Query_user()
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(q))
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
	data, _ := ioutil.ReadAll(response.Body)

	log.Println(string(data))

	data_resp := Request{}

	if err := json.Unmarshal(data, &data_resp); err != nil {
		log.Println("Can not unmnarshal JSON")
	}
	log.Println("This -> HTTP Response: %s")
	log.Println(data_resp.Data.Loremsite_Users[0])

	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func SignupSubmit(c *fiber.Ctx) error {
	//*****/ Query User

	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
	return nil
  
}

func ReadGen(c *fiber.Ctx) error { 

	return c.Render("readgenerator", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderGenerators(c *fiber.Ctx) error {
	//*********/ Query User
	gens := fetchGenerators()

	return c.Render("listgens", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
		"Loremgens":  gens,
	}, "layouts/htm")
}

func RenderContact(c *fiber.Ctx) error {
	//*********/ Query User

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}
