package main

// import (
// 	"bytes"
// 	"context"
// 	"fmt"

// 	"github.com/codeWithUtkarsh/go-abs/abs"
// 	nft "github.com/codeWithUtkarsh/go-abs/nftstorage"
// )

// func main() {
// 	jsonData := []byte(`{"name": "John", "age": 30}`)
// 	reader := bytes.NewReader(jsonData)

// 	client, _ := nft.NewClient("nft", nft.WithToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDk3ODMwMGIzNUMzOEM3NzMxYWNDNjk4NDFiODU4NDNiRmIxRjExN0QiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY0ODY2MjM1MDcwNiwibmFtZSI6Im9ydGVsaXVzX3Nib21fbGVkZ2VyX2RlbW8ifQ.egQzHqDB83UoK7ynE4fn4dhgIKWLhPP1B9E6qey4KHM"))
// 	clientHub := abs.NewClientHub(client)

// 	uploadParam := abs.UploadParam{IOReader: reader}
// 	uploadResults, uploadErrors := clientHub.Upload(context.Background(), uploadParam)
// 	fmt.Println("Upload Results:", uploadResults)
// 	fmt.Println("Upload Errors:", uploadErrors)

// 	status, err := client.Status(context.Background(), "bafkreihn7cjn5gxhtdwp4ufjqdpozqxlhtcgv6z5ww5swivq22ix46x7uy")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Created:", status)
// }

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codeWithUtkarsh/go-abs/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// dbFactory := DatabaseFactory{}
	// datasource, err = dbFactory.CreateDatabase()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db, err = datasource.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/upload", handler.Create).Methods("POST")
	myRouter.HandleFunc("/status", handler.GetStatus).Methods("GET")

	handler.SetSwagger(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
