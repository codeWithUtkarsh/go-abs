package abs

import (
	"context"
	"io"
)

type Metadata = map[string]interface{}

type Cid = string
type PinStatus = string

type Client interface {
	Upload(context.Context, Payload) (Metadata, error)
	Status(context.Context, Cid) (Metadata, error)
}

type Payload struct {
	IOReader io.Reader `json:"IOReader"`
}

type Result struct {
	Metadata map[string]interface{} `json:"metadata"`
}
