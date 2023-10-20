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
		Get("/api/session/data/{data-source}/activeConnections")

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
		Patch("/api/session/data/{data-source}/activeConnections")

	return
}

func (ac *ActiveConnection) Kills(identifiers []string) (err error) {
	operations := make([]vars.Operation, 0)

	for _, id := range identifiers {
		operations = append(operations, vars.Operation{
			OP:   "remove",
			Path: fmt.Sprintf("/%s", id),
		})
	}

	_, err = ac.NewRequest().
		SetBody(operations).
		Patch("/api/session/data/{data-source}/activeConnections")

	return
}
