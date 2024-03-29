package ql

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestQL(t *testing.T) {
	t.Run("example-1", func(t *testing.T) {
		// Schema
		fields := graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			},
		}
		rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
		schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
		schema, err := graphql.NewSchema(schemaConfig)
		if err != nil {
			log.Fatalf("failed to create new schema, error: %v", err)
		}

		// Query
		query := `
        {
            hello
        }
    `
		params := graphql.Params{Schema: schema, RequestString: query}
		r := graphql.Do(params)
		if len(r.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		}
		rJSON, _ := json.Marshal(r)
		fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
	})
}
