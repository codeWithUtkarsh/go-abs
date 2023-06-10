package nftstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (cli *client) Status(ctx context.Context, cid string) (pinStatus string, err error) {

	url := cli.conf.Endpoint + "/check/" + cid
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New("failed to create HTTP request")
	}

	req.Header.Set("Authorization", "Bearer "+cli.conf.AccessToken)
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to send HTTP request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("[%d] %s", resp.StatusCode, string(body))
	}

	var response NftSuccessResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", errors.New("failed to unmarshal response JSON")
	}

	return response.Value.Pin.CID, nil
}
