package model

import (
	"context"

	"github.com/thanhhm/stw/util"
)

type Caller interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type clImpl struct {
	hc *util.HTTPC
}

func NewCaller(hc *util.HTTPC) Caller {
	return &clImpl{
		hc: hc,
	}
}

func (cl *clImpl) Get(ctx context.Context, url string) ([]byte, error) {
	req := util.HTTPCRequest{
		Method: "GET",
		Path:   url,
	}

	resp, err := cl.hc.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.RespBody, nil
}
