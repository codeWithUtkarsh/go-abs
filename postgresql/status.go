package postgresql

import (
	"context"
)

func (cli *client) Status(ctx context.Context, cid string) (pinStatus map[string]interface{}, err error) {

	response := make(map[string]interface{})
	response["message"] = "This API is Not Availaible for Relational Databases"
	return response, nil
}
