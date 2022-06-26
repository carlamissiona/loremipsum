package controllers

import (
   "loremipsumbytes/pkg/hasura"
  "net/http"
  "github.com/gofiber/fiber/v2"
  "io/ioutil"
  "log"
  	"encoding/json"
    "time"
  "fmt"
  "bytes" 
  )


var Url = "https://loremdb.hasura.app/v1/graphql"

func RenderIndex(c *fiber.Ctx) error {
 //*********/ Query Users 
  q := hasura.Query_user()

  req , err := http.NewRequest("POST", Url, bytes.NewBuffer(q)) 
  if err != nil {
    log.Println("Error in Gql Run")
    
  }
  client := &http.Client{Timeout: time.Second * 7}
  response, err := client.Do(req)
  defer response.Body.Close()
   if err != nil {
    log.Println("Error in Clieee snt Do Run")
    }
  data, _ := ioutil.ReadAll(response.Body)
  log.Println(string(data)) ; log.Println("%#v\n, data");   
    log.Println("%#v\n, data") ;   log.Println("%#v\n, data");   log.Println("%#v\n, data")
   fmt.Printf("======%#v\n", response)
  log.Println("===%#v\n, data")
  log.Println("===%#v\n, data\n\n\n\n\n")
  log.Println("RESPONSE HERE===%#v\n, data\n\n\n\n\n")
  var jrsp map[string]interface{} = nil

  if err := json.Unmarshal(data, &jrsp); err != nil {  
        log.Println("Can not unmnarshal JSON")
    }
  log.Printf("",response) 
  log.Println("Can  unmarshal JSON")
   log.Println("!an  unmarshal JSON")
     log.Println("!!!!!Can  unmarshal JSON")
     log.Println("Can  unmarshal JSON")
   log.Println(jrsp) 
  
json.NewDecoder(response.Body).Decode(jrsp)
log.Println("rense struct:", response)  
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	},"layouts/htm")
}     
  