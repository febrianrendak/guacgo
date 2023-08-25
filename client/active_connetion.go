package client

import (
	"fmt"
	"github.com/febrianrendak/guacgo/vars"
)

type ActiveConnection struct {
	*Client
}

func (client *Client) ActiveConnection() *ActiveConnection {
	return &ActiveConnection{
		client,
	}
}

func (ac *ActiveConnection) List() (activeConnectionList []vars.ActiveConnection, mapOfActiveConnection map[string]vars.ActiveConnection, err error) {
	_, err = ac.NewRequest().
		SetSuccessResult(&mapOfActiveConnection).
		Get("/guacamole/api/session/data/{data-source}/activeConnections")

	for _, userGroup := range mapOfActiveConnection {
		activeConnectionList = append(activeConnectionList, userGroup)
	}

	return
}

func (ac *ActiveConnection) Kill(identifier string) (err error) {
	_, err = ac.NewRequest().
		SetBody([]vars.Operation{
			{
				OP:   "remove",
				Path: fmt.Sprintf("/%s", identifier),
			},
		}).
		Patch("/guacamole/api/session/data/{data-source}/activeConnections")

	return
}
