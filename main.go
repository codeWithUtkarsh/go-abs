package main

import (
	"bytes"
	"context"
	"fmt"

	nft "github.com/codeWithUtkarsh/go-abs/nftstorage/"
)

func main() {
	jsonData := []byte(`{"name": "John", "age": 30}`)
	reader := bytes.NewReader(jsonData)

	client := nft.NewClient("nft", nft.WithEndpoint)
	clientHub := NewClientHub(client)

	uploadParam := UploadParam{IOReader: reader}
	uploadResults, uploadErrors := clientHub.Upload(context.Background(), uploadParam)
	fmt.Println("Upload Results:", uploadResults)
	fmt.Println("Upload Errors:", uploadErrors)

	// status, err := client.Status(context.TODO(), cid)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Created:", status)
}
