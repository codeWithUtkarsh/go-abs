package nftstorage

const endpoint = "https://api.nft.storage"

type clientConfig struct {
	mimes       []string
	endpoint    string
	accesstoken string
}

type Option func(*clientConfig) error

func WithEndpoint(endpoint string) Option {
	return func(nsc *clientConfig) error {
		nsc.endpoint = endpoint
		return nil
	}
}

func WithToken(accessToken string) Option {
	return func(nsc *clientConfig) error {
		nsc.accesstoken = accessToken
		return nil
	}
}

func WithMimes(mimes ...string) Option {
	return func(nsc *clientConfig) error {
		nsc.mimes = mimes
		return nil
	}
}
