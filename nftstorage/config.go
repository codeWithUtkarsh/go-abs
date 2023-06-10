package nftstorage

const defaultEndpoint = "https://api.nft.storage"

type ClientConfig struct {
	Mimes       []string
	Endpoint    string
	AccessToken string
}

type Option func(*ClientConfig) error

func WithEndpoint(endpoint string) Option {
	return func(cfg *ClientConfig) error {
		cfg.Endpoint = endpoint
		return nil
	}
}

func WithToken(accessToken string) Option {
	return func(cfg *ClientConfig) error {
		cfg.AccessToken = accessToken
		return nil
	}
}

func WithMimes(mimes ...string) Option {
	return func(cfg *ClientConfig) error {
		cfg.Mimes = mimes
		return nil
	}
}
