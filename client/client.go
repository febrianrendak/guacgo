package client

import (
	"fmt"
	"github.com/febrianrendak/guacgo/vars"
	"github.com/imroc/req/v3"
)

type Client struct {
	c           *req.Client
	username    string
	password    string
	apiEndpoint string
	tokenData   vars.TokenResp
}

func NewClient(apiEndpoint, username, password string) *Client {
	client := req.C().
		SetCommonErrorResult(&vars.ErrorResp{}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return nil
			}

			if errMsg, ok := resp.ErrorResult().(*vars.ErrorResp); ok {
				resp.Err = errMsg
				return nil
			}

			if !resp.IsSuccessState() {
				resp.Err = fmt.Errorf("status code: %d, err: %s", resp.StatusCode, resp.Err)
			}
			return nil
		})
	client.BaseURL = apiEndpoint

	return &Client{
		apiEndpoint: apiEndpoint,
		username:    username,
		password:    password,
		c:           client,
		tokenData:   vars.TokenResp{},
	}
}

func (client *Client) NewRequest() *req.Request {
	return client.c.R().
		SetHeader("Guacamole-Token", client.tokenData.AuthToken).
		SetPathParam("data-source", client.tokenData.DataSource).
		SetRetryCount(1).
		AddRetryHook(func(resp *req.Response, err error) {
			tokenResp, err := client.Auth().Token()
			if err == nil {
				client.tokenData = tokenResp

				resp.Request.SetHeader("Guacamole-Token", tokenResp.AuthToken)
				resp.Request.SetPathParam("data-source", tokenResp.DataSource)
			}
		}).
		AddRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode == 403
		})
}
