package controllers

import (
    "loremipsumbytes/pkg/hasura"
    "net/http"
    "github.com/gofiber/fiber/v2"
    "io/ioutil"
    "log"
    "encoding/json"
    "time"
    // "fmt"
    "bytes"
) 

 type DataForm struct {
    Loremsite_Users   []User     `json:"loremsite_users"` 
} 
type User struct {
    Full_name string   `json:"full_name"`
    Email string    `json:"email"`
} 
type Request struct {
    Data   DataForm 
}   
var Url = "https://loremdb.hasura.app/v1/graphql"
 
func RenderIndex(c * fiber.Ctx) error {
    //**********/ Query User    
    q:= hasura.Query_user()
    req, err:= http.NewRequest("POST", Url, bytes.NewBuffer(q))
    if err != nil {
        log.Println("Error in Gql Run") 
    }
  
    client:= &http.Client {
        Timeout: time.Second*7,
    }
    response, err:= client.Do(req)
    defer response.Body.Close()
    if err != nil {
        log.Println("Error in Clieee snt Do Run")
    }
    data_bytes, _:= ioutil.ReadAll(response.Body)
    
    log.Println(string(data_bytes));
    
    data_resp := Request{}
    err = json.Unmarshal(data_bytes, &data_resp) 
     
    if err != nil  {
        log.Println("Can not unmnarshal JSON")
    } 
    log.Println("This -> HTTP Response: %s"); 
    log.Println(data_resp.Data.Loremsite_Users[0])
     
    return c.Render("index", fiber.Map {
        "FiberTitle": "Hello From Fiber Html Engine",
    }, "layouts/htm")
}
          
 
func RenderAddGraph(c * fiber.Ctx) error {
    //**********/ Query User    
    q:= hasura.Query_user()
    req, err:= http.NewRequest("POST", Url, bytes.NewBuffer(q))
    if err != nil {
        log.Println("Error in Gql Run") 
    }






         
    client:= &http.Client {
        Timeout: time.Second*7,
    }
    response, err:= client.Do(req)
    defer response.Body.Close()
    if err != nil {
        log.Println("Error in Clieee snt Do Run")
    }
    data, _:= ioutil.ReadAll(response.Body)
    
    log.Println(string(data));
    
    data_resp := Request{}
     
    if err:= json.Unmarshal(data, &data_resp) ; err != nil {
        log.Println("Can not unmnarshal JSON")
    } 
    log.Println("This -> HTTP Response: %s"); log.Println(data_resp.Data.Loremsite_Users[0])
     
    return c.Render("index", fiber.Map {
        "FiberTitle": "Hello From Fiber Html Engine",
    }, "layouts/htm")
}
     