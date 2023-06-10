package nftstorage

import (
	"errors"

	main "github.com/codeWithUtkarsh/go-abs"
)

const (
	MAX_REQUEST_BODY_SIZE = 1024 * 1024 * 99 // 99 MiB
)

var ErrRequestBodyLimit = errors.New("body size max limit 99MiB")

type client struct {
	name string
	conf clientConfig
}

func (c *client) Name() string {
	return c.name
}

func NewClient(name string, opts ...Option) (*client, error) {
	conf := clientConfig{
		endpoint: endpoint,
	}

	for _, opt := range opts {
		if err := opt(&conf); err != nil {
			return nil, err
		}
	}

	cli := client{name: name, conf: conf}

	return &cli, nil
}

var _ main.Client = (*client)(nil)
