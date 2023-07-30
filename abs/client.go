package abs

import (
	"context"
	"fmt"
)

type ClientHub struct {
	clis []Client
}

func NewClientHub(clis ...Client) *ClientHub {
	cli := ClientHub{clis: clis}
	return &cli
}

func (c *ClientHub) Upload(ctx context.Context, file Payload) (ress []Result, errs []error) {
	fmt.Println(file)
	fmt.Println(c)
	fmt.Println(c.clis)
	res, err := c.clis[0].Upload(ctx, file)
	if err != nil {
		errs = append(errs, err)
	}
	ress = append(ress, Result{Metadata: res})

	return ress, errs
}

func (c *ClientHub) Status(ctx context.Context, cid Cid) (ress []Result, errs []error) {
	res, err := c.clis[0].Status(ctx, cid)
	if err != nil {
		errs = append(errs, err)
	}
	ress = append(ress, Result{Metadata: res})

	return
}
