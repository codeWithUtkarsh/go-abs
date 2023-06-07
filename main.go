package ipfsstorage

import (
	"bytes"
	"context"
	"fmt"
)

func main() {
	jsonData := []byte(`{"name": "John", "age": 30}`)
	reader := bytes.NewReader(jsonData)

	client := &ClientHub{}

	up := UploadParam{IOReader: reader}
	cid, err := client.Upload(context.TODO(), up)
	if err != nil {
		panic(err)
	}
	fmt.Println("CID:", cid)

	// status, err := client.Status(context.TODO(), cid)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Created:", status)
}
