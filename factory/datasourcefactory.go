package factory

import (
	"errors"
	"fmt"

	"github.com/codeWithUtkarsh/go-abs/abs"
	"github.com/codeWithUtkarsh/go-abs/nftstorage"
)

type DataSourceFactory struct{}

func (df DataSourceFactory) CreateDatasourceClient() (abs.Client, error) {

	datasource := getenv("DATA_SOURCE", "")

	switch datasource {
	case "NFT":
		API_KEY := getenv("API_KEY", "")
		client, err := nftstorage.NewClient("NFT", nftstorage.WithToken(API_KEY))
		if err != nil {
			return nil, errors.New("failed to create datasource in factory")
		}
		return client, nil
	// case "mysql":
	// 	port, _ := strconv.Atoi(getenv("DB_PORT", "3306"))
	// 	return MySQL{
	// 		host:     getenv("DB_HOST", "localhost"),
	// 		port:     port,
	// 		username: getenv("DB_USERNAME", "root"),
	// 		password: getenv("DB_PASSWORD", "password"),
	// 		dbname:   getenv("DB_NAME", "testdb"),
	// 	}, nil
	// case "postgres":
	// 	port, _ := strconv.Atoi(getenv("DB_PORT", "3306"))
	// 	return PostgreSQL{
	// 		host:     getenv("DB_HOST", "localhost"),
	// 		port:     port,
	// 		username: getenv("DB_USERNAME", "root"),
	// 		password: getenv("DB_PASSWORD", "password"),
	// 		dbname:   getenv("DB_NAME", "testdb"),
	// 	}, nil
	default:
		return nil, fmt.Errorf("unsupported database type")
	}
}
