package util

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type HTTPC struct {
	client *http.Client
}

type HTTPCConfig struct {
	Timeout time.Duration
}

func NewHTTPC(cfg *HTTPCConfig) *HTTPC {
	return &HTTPC{&http.Client{Timeout: cfg.Timeout}}
}

type HTTPCRequest struct {
	Path   string
	Method string
	Body   interface{}
}

type HTTPCResponse struct {
	Code     int64
	RespBody []byte
}

func (hc *HTTPC) DoAndUnmarshal(ctx context.Context, data HTTPCRequest, v interface{}) (*HTTPCResponse, error) {
	httpcResp, err := hc.Do(ctx, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(httpcResp.RespBody, v)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal response error")
	}

	return httpcResp, nil
}

func (hc *HTTPC) Do(ctx context.Context, data HTTPCRequest) (*HTTPCResponse, error) {
	var reqBody io.Reader
	if data.Body != nil {
		b, err := json.Marshal(data.Body)
		if err != nil {
			return nil, errors.Wrap(err, "Marshal request error")
		}
		reqBody = bytes.NewBuffer(b)
	}
	req, err := http.NewRequestWithContext(ctx, data.Method, data.Path, reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := HTTPCResponse{
		Code:     int64(resp.StatusCode),
		RespBody: respBody,
	}
	return &result, err
}
