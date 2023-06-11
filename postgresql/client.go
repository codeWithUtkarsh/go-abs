package postgresql

import (
	"database/sql"
	"errors"
	"fmt"

	abs "github.com/codeWithUtkarsh/go-abs/abs"
)

type client struct {
	host       string
	port       int
	username   string
	password   string
	dbname     string
	connection *sql.DB
}

func (c *client) Name() string {
	return c.dbname
}

func NewClient(
	host string, port int, username string, password string, dbname string) (*client, error) {

	cli := &client{
		host:       host,
		port:       port,
		username:   username,
		password:   password,
		dbname:     dbname,
		connection: nil,
	}

	err := Connect(cli)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func Connect(cli *client) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cli.host,
		cli.port,
		cli.username,
		cli.password,
		cli.dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return errors.New("connection to postgres failed")
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	cli.connection = db
	return nil
}

var _ abs.Client = (*client)(nil)
