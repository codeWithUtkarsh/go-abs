package ipfsstorage

import (
	"context"
)

type ClientHub struct {
	cli Client
}

func NewClientHub(cli Client) *Client {
	return &cli
}

func (c *ClientHub) Upload(ctx context.Context, file UploadParam) (ress []UploadResult, errs []error) {
	res, err := c.cli.Upload(ctx, file)
	if err != nil {
		errs = append(errs, err)
	}
	ress = append(ress, UploadResult{CID: res})

	return ress, errs
}

func (c *ClientHub) Status(ctx context.Context, cid Cid) (ress []StatusResult, errs []error) {
	res, err := c.cli.Status(ctx, cid)
	if err != nil {
		errs = append(errs, err)
	}
	ress = append(ress, StatusResult{Status: res})

	return
}
