package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codeWithUtkarsh/go-abs/abs"
	"github.com/codeWithUtkarsh/go-abs/factory"
	model "github.com/codeWithUtkarsh/go-abs/handler/models"
	_ "github.com/go-sql-driver/mysql"
	swg "github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-http/swagger"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Create(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	jsonData := []byte(body)
	reader := bytes.NewReader(jsonData)

	dFactory := factory.DataSourceFactory{}
	client, _ := dFactory.CreateDatasourceClient()

	clientHub := abs.NewClientHub(client)
	payload := abs.Payload{IOReader: reader}
	response, uploadErrors := clientHub.Upload(context.Background(), payload)

	if len(uploadErrors) != 0 {
		panic(uploadErrors[0].Error())
	}
	fmt.Println("Upload Results:", response)

	result := model.GenericResponse{Result: "success", Metadata: response[0].Metadata}
	json.NewEncoder(w).Encode(result)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// _, err := validate()
	// if err != nil {
	// 	panic(err.Error())
	// }
	dFactory := factory.DataSourceFactory{}
	client, _ := dFactory.CreateDatasourceClient()
	clientHub := abs.NewClientHub(client)

	response, err := clientHub.Status(context.Background(), "bafkreihn7cjn5gxhtdwp4ufjqdpozqxlhtcgv6z5ww5swivq22ix46x7uy")
	if err != nil {
		panic(err)
	}

	result := model.GenericResponse{Result: "success", Metadata: response[0].Metadata}
	json.NewEncoder(w).Encode(result)
}

func SetSwagger(myRouter *mux.Router) {
	endpoints := []swg.Endpoint{
		swg.EndPoint(swg.POST, "/upload", "Upload", swg.Params(), model.RequestPayload{}, model.GenericResponse{}, nil, "", nil),
		swg.EndPoint(swg.GET, "/status", "Status", swg.Params(swg.StrParam("cid", true, "")), nil, model.GenericResponse{}, nil, "Get status of CID", nil),
	}

	sw := swg.CreateNewSwagger("Swagger API", "1.0")
	swg.AddEndpoints(endpoints)

	myRouter.PathPrefix("/swagger").Handler(swagger.SwaggerHandler(sw.GenerateDocs()))
	// if you want to export your swagger definition to a file
	// sw.ExportSwaggerDocs("docs/swagger.json") // optional
}
