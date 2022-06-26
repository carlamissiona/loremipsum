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
    "github.com/davecgh/go-spew/spew"
)
 

type User struct {
    Full_name string
    Email string
}

type Response struct {
    Msg    string
    User   interface{}
    Status int
    Data       []Users
}
// var post Article 
//   json.Unmarshal(reqBody, &post)
var Url = "https://loremdb.hasura.app/v1/graphql"

func RenderIndex(c * fiber.Ctx) error {
    //*********/ Query Users 
    q:= hasura.Query_user()
    req, err:= http.NewRequest("POST", Url, bytes.NewBuffer(q))
    if err != nil {
        log.Println("Error in Gql Run") }
  
    client:= & http.Client {
        Timeout: time.Second * 7
    }
    response, err:= client.Do(req)
    defer response.Body.Close()
    if err != nil {
        log.Println("Error in Clieee snt Do Run")
    }
    data, _:= ioutil.ReadAll(response.Body)
     log.Println("string(data)");  log.Println("string(data)");  log.Println("string(data)");

    log.Println(string(data));
    log.Println("%#v\n, data");
    log.Println("%#v\n, data");log.Println("%#v\n, data");log.Println("%#v\n, data")
    fmt.Printf("======%#v\n", response.Body)

    var jrsp map[string]interface {} = nil

    if err:= json.Unmarshal(data, & jrsp);err != nil {
        log.Println("Can not unmnarshal JSON")
    }
    log.Printf("", response)

    log.Println(jrsp)
    fmt.Printf("---\ngo-spew-----------\n")
    spew.Dump(jrsp)
    json.NewDecoder(response.Body).Decode(jrsp)
    log.Println(" structaaa:", jrsp)
    return c.Render("index", fiber.Map {
        "FiberTitle": "Hello From Fiber Html Engine",
    }, "layouts/htm")
}
 
// type Response struct {
//     Data       []Users
   
// }

// type Users struct {
//     Data string      `json:"data"`
//     Key string            `json:"key"`
//     Value string          `json:"value"`
// }

// func main() {
//     s := string(`{"operation": "get", "key": "example"}`)
//     data := Request{}
//     json.Unmarshal([]byte(s), &data)
//     fmt.Printf("Operation: %s", data.Operation)
// }
// /////////

// type Request struct {
//     Operation string      `json:"operation"`
//     Key string            `json:"key"`
//     Value string          `json:"value"`
// }

// func main() {
//     s := string(`{"operation": "get", "key": "example"}`)
//     data := Request{}
//     json.Unmarshal([]byte(s), &data)
//     fmt.Printf("Operation: %s", data.Operation)
// }

// ..
// import (
//   "encoding/json"
//   "fmt"
//   "io/ioutil"
//   "log"
//   "net/http"

//   "github.com/gorilla/mux"
// )

// type Article struct {
//   Id string `json:"Id"`
//   Title string `json:"Title"`
//   Content string `json:"Content"`
//   Summary string `json:"Summary"`
// }

// ...

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
//   reqBody, _ := ioutil.ReadAll(r.Body)
//   var post Article 
//   json.Unmarshal(reqBody, &post)

//   json.NewEncoder(w).Encode(post)

//   newData, err := json.Marshal(post)
//   if err != nil {
//     fmt.Println(err)
//   } else {
//     fmt.Println(string(newData))
//   }
// }

// func handleReqs() {
//   r := mux.NewRouter().StrictSlash(true)
//   r.HandleFunc("/post", createNewArticle).Methods("POST")

//   log.Fatal(http.ListenAndServe(":8000", r))
// }

// func main() {
//   handleReqs();
// }