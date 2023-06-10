package nftstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	ipfsstorage "github.com/codeWithUtkarsh/go-abs"
)

func (cli *client) Status(ctx context.Context, cid string) (pinStatus ipfsstorage.PinStatus, err error) {

	url, _ := url.Parse(cli.conf.endpoint + "/check/" + cid)

	req := http.Request{
		URL: url,
		Header: http.Header{
			"Authorization": {"Bearer " + cli.conf.accesstoken},
			// "Content-Type":  {"application/car"},
			"Accept": {"application/json"},
		},
		Method: http.MethodGet,
	}

	httpCli := http.Client{}

	response, err := httpCli.Do(&req)
	if err != nil {
		err = errors.New("http request error")
		return
	}
	defer response.Body.Close()

	resBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New("ioutile read body error")
		return
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("[%d]%s", response.StatusCode, string(resBytes)))
		return
	}

	var res Response200
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		err = errors.New("json unmarshal response error")
		return
	}

	return res.Value.Pin.Status, nil
}
