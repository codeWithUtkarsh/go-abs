package impfstorage

import (
	"context"
	"io"
)

type Cid = string
type PinStatus = string

type Client interface {
	Upload(context.Context, UploadParam) (Cid, error)
	Status(context.Context, Cid) (PinStatus, error)
}

type UploadParam struct {
	IOReader io.Reader `json:"IOReader"`
}

type UploadResult struct {
	CID string `json:"cid"`
}

type StatusResult struct {
	Status PinStatus `json:"status"`
}
