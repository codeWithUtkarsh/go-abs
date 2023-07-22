package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	abs "github.com/codeWithUtkarsh/go-abs/abs"
	model "github.com/codeWithUtkarsh/go-abs/handler/models"
	_ "github.com/go-sql-driver/mysql"
)

func (cli *client) Upload(ctx context.Context, payload abs.Payload) (metadata map[string]interface{}, err error) {

	response := make(map[string]interface{})

	var p model.RequestPayload
	err = json.NewDecoder(payload.IOReader).Decode(&p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	entity := p.Entity
	persistantData := p.Payload

	fmt.Println(entity)
	obj, exists := onboardedEntitiesMap[entity]
	if !exists {
		fmt.Println(entity)
		fmt.Println("Invalid Entity type: You Need to onboard this entity definition by updation models.go in ./postgresql")
		return
	}

	queryf := `INSERT INTO public.` + entity + ` (`
	values := `VALUES (`

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		key := strings.ToLower(field.Name)
		value, ok := persistantData[key]
		if ok {
			fmt.Printf("Key '%s' exists with value: %d\n", key, value)
			queryf += key + `,`
			values += `'` + value.(string) + `',`
		}

		fmt.Println("Field Name:", field.Name)
	}
	queryf = queryf[:len(queryf)-1]
	values = values[:len(values)-1]

	values += `) `
	queryf += `) ` + values + ` RETURNING id`

	fmt.Println(queryf)
	id := 0
	err = cli.connection.QueryRow(queryf).Scan(&id)
	if err != nil {
		return response, err
	}

	response["message"] = "The User has been inserted successfully! Id: " + fmt.Sprint(id)
	return response, nil
}
