package impfstorage

import (
	"context"
)

type ClientHub struct {
	clis []Client
}

func NewClientHub(clis ...Client) *ClientHub {
	cli := ClientHub{clis: clis}
	return &cli
}

func (c *ClientHub) Upload(ctx context.Context, file UploadParam) (ress []UploadResult, errs []error) {
	res, err := c.clis[0].Upload(ctx, file)
	if err != nil {
		errs = append(errs, err)
	}
	ress = append(ress, UploadResult{CID: res})

	return ress, errs
}

func (c *ClientHub) Status(ctx context.Context, cid Cid) (ress []StatusResult, errs []error) {
	res, err := c.clis[0].Status(ctx, cid)
	if err != nil {
		errs = append(errs, err)
	}
	ress = append(ress, StatusResult{Status: res})

	return
}
