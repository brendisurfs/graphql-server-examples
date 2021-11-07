package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user
/*
	Create user object type with fields "id" and "name" by using GraphqlObjectTypeConfig.
*/

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

//now create a query object type with fields "user" has type [UserType]
//QUERY TYPE HERE
// |
// v
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams)(interface{}, error){

					//I have to figure out this .(string) thing.
					idQuery, OK := p.Args["id"].(string)
					if OK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Description: "return all users in the json file",
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error){
					return data, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	})


func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params {
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, expected errors: %v", result.Errors)
	}
	return result
}

func importJSONFromFile(filename string, result interface{}) (isOk bool){
	isOk = true
	content, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		fmt.Println("Error: ", readErr)
		isOk = false
	}
	readErr = json.Unmarshal(content, result)
	if readErr != nil {
		isOk = false
		println("error: ", readErr)
	}
	return
}

func main() {
	//import data setup
	_ = importJSONFromFile("data.json", &data)
	//	graphql setup

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		//alright, so we are braking down the url params, because gql throws the query in the url.
		result := executeQuery(r.URL.Query().Get("query"), schema)

		//writes to w.
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("server now running on port 8080")

	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	//server start
	http.ListenAndServe(":8080", nil)
}
