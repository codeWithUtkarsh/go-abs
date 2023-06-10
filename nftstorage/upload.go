package nftstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	abs "github.com/codeWithUtkarsh/go-abs/abs"
)

func (cli *client) Upload(ctx context.Context, file abs.UploadParam) (cid string, err error) {

	url := cli.conf.Endpoint + "/upload"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, file.IOReader)
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

	fmt.Printf("%s", string(body))

	return response.Value.CID, nil
}
