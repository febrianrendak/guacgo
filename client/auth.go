package client

import "github.com/febrianrendak/guacgo/vars"

type Auth struct {
	*Client
}

func (client *Client) Auth() *Auth {
	return &Auth{
		client,
	}
}

func (auth *Auth) Token() (vars.TokenResp, error) {
	respJson := vars.TokenResp{}
	_, err := auth.c.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"username": auth.username,
			"password": auth.password,
		}).
		SetSuccessResult(&respJson).
		Post("/api/tokens")

	return respJson, err
}
