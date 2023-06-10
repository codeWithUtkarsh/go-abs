package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codeWithUtkarsh/go-abs/abs"
	model "github.com/codeWithUtkarsh/go-abs/handler/models"
	nft "github.com/codeWithUtkarsh/go-abs/nftstorage"
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

	client, _ := nft.NewClient("nft", nft.WithToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDk3ODMwMGIzNUMzOEM3NzMxYWNDNjk4NDFiODU4NDNiRmIxRjExN0QiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY0ODY2MjM1MDcwNiwibmFtZSI6Im9ydGVsaXVzX3Nib21fbGVkZ2VyX2RlbW8ifQ.egQzHqDB83UoK7ynE4fn4dhgIKWLhPP1B9E6qey4KHM"))
	clientHub := abs.NewClientHub(client)

	uploadParam := abs.Payload{IOReader: reader}
	response, uploadErrors := clientHub.Upload(context.Background(), uploadParam)

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
	client, _ := nft.NewClient("nft",
		nft.WithEndpoint("https://api.nft.storage"),
		nft.WithToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDk3ODMwMGIzNUMzOEM3NzMxYWNDNjk4NDFiODU4NDNiRmIxRjExN0QiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY0ODY2MjM1MDcwNiwibmFtZSI6Im9ydGVsaXVzX3Nib21fbGVkZ2VyX2RlbW8ifQ.egQzHqDB83UoK7ynE4fn4dhgIKWLhPP1B9E6qey4KHM"))
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
