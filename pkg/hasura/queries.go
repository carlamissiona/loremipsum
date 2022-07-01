package hasura

import (
	"encoding/json"

	"log"
)

func Query_gens() []byte {

	gr := map[string]string{
		"query": ` query MyQuery {
              loremsite_generators(limit: 3) {
                link
                id
                name
                desc
              }
            }

    `,
	}
	log.Println("============>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal")
	}
	log.Println(" Query_gens")
	log.Println(" Query_gens")
	log.Println(" Query_gens")
	return json
}

func Query_user_asc() []byte {

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
	log.Println("========>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal")
	}

	// log.Println(string(json))
	// log.Println("string(json)")
	return json
}

func Query_gens_byname() []byte {

	gr := map[string]string{
		"query": ` query MyQuery {
          loremsite_users(limit: 10){
            email
            full_name
            id
          }
        }
    `,
	}
	log.Println("============>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal")
	}

	return json
}

func Query_user() []byte {

	gr := map[string]string{
		"query": ` query MyQuery {
          loremsite_users(limit: 3) {
            email
            full_name
            id
          }
        }
    `,
	}
	log.Println("============>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal")
	}

	return json
}

func Mutation_add_gen(desc string, link string, name string) []byte {
	log.Println(desc)
	log.Println(link)
	gq := map[string]string{
		"query": `
      mutation MyMutation {
        insert_loremsite_generators_one(object: {desc: "` + desc + `", link: "` + link + `", name: "` + name + `"})
      }
    
    `,
	}

	json, err := json.Marshal(gq)
	if err != nil {
		log.Println("Error Marshal")
	}

	return json

}

func Mutation_signup_user(email string, fullname string) []byte {
	log.Println(email)
	log.Println(email)
	gr := map[string]string{
		"query": `mutation MyMutation {
      insert_loremsite_users_one(object: {full_name: "` + fullname + `", password: "ewew", email: "` + email + `"}) {
        id
      }
    }
    `,
	}

	log.Println(">>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal")
	}

	return json
}
