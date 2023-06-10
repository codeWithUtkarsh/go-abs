package nftstorage

import (
	abs "github.com/codeWithUtkarsh/go-abs/abs"
)

type client struct {
	name string
	conf ClientConfig
}

func (c *client) Name() string {
	return c.name
}

func NewClient(name string, opts ...Option) (*client, error) {
	conf := ClientConfig{
		Endpoint: defaultEndpoint,
	}

	for _, opt := range opts {
		if err := opt(&conf); err != nil {
			return nil, err
		}
	}

	cli := client{name: name, conf: conf}

	return &cli, nil
}

var _ abs.Client = (*client)(nil)
