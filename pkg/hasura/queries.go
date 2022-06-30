package hasura
import (
	"encoding/json"
	 
	"log" 

 )


 func Query_gens() []byte  {
 
  gr := map[string]string{
    "query": ` query MyQuery {
              loremsite_generators {
                link
                id
                name
                desc
              }
            }

    `,
  }
  log.Println("============>>>gr"); log.Println(gr)
  json , err := json.Marshal(gr)
   if err != nil {
     log.Println("Error Marshal")
   }
   log.Println(" Query_gens") ;     log.Println(" Query_gens") ;    log.Println(" Query_gens")
  return json
} 

  




 func Query_user_asc() []byte  {
 
  gr := map[string]string{
    "query": ` query MyQuery {
          loremsite_users(order_by: {full_name: asc}) {
            email
            full_name
            id
          }
        }
    `,
  }
  log.Println("========>>>gr"); log.Println(gr)
  json , err := json.Marshal(gr)
   if err != nil {
     log.Println("Error Marshal")
   }
   
  // log.Println(string(json))
  // log.Println("string(json)")
  return json
} 

  
 
func Query_gens_byname() []byte  {
 
  gr := map[string]string{
    "query": ` query MyQuery {
          loremsite_users {
            email
            full_name
            id
          }
        }
    `,
  }
  log.Println("============>>>gr"); log.Println(gr)
  json , err := json.Marshal(gr)
   if err != nil {
     log.Println("Error Marshal")
   }
    
  return json
} 


 


 func Query_user() []byte  {
 
  gr := map[string]string{
    "query": ` query MyQuery {
          loremsite_users {
            email
            full_name
            id
          }
        }
    `,
  }
  log.Println("============>>>gr"); log.Println(gr)
  json , err := json.Marshal(gr)
   if err != nil {
     log.Println("Error Marshal")
   }
    
  return json
} 



 func Mutation_signup_user(email string , fullname string) []byte  {
 
  gr := map[string]string{
    "query": ` 
      mutation MyMutation {
        insert_loremsite_users_one(objects: {email: "`+ email +`", full_name: "`+ fullname + `"})
      }

    `,
  }
   
  log.Println(">>gr"); log.Println(gr)
  json , err := json.Marshal(gr)
   if err != nil {
     log.Println("Error Marshal")
   }
    
  return json
}  