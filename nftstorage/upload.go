package nftstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	ipfsstorage "github.com/codeWithUtkarsh/go-abs"
)

func (cli *client) Upload(ctx context.Context, file ipfsstorage.UploadParam) (cid string, err error) {

	url, _ := url.Parse(cli.conf.endpoint + "/upload")

	req := http.Request{
		URL: url,
		Header: http.Header{
			"Authorization": {"Bearer " + cli.conf.accesstoken},
			// "Content-Type":  {"application/car"},
			"Accept": {"application/json"},
		},
		Method: http.MethodPost,
		Body:   io.NopCloser(file.IOReader),
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

	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("[%d]%s", response.StatusCode, string(resBytes)))
		return
	}

	var res Response200
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		err = errors.New("json unmarshal response error")
		return
	}

	fmt.Printf("%s", string(resBytes))

	return res.Value.CID, nil
}
