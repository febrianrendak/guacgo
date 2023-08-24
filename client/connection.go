package client

import (
	"github.com/febrianrendak/guacgo/vars"
)

type Connection struct {
	*Client
}

func (client *Client) Connection() *Connection {
	return &Connection{
		client,
	}
}

func (connection *Connection) List() ([]vars.Connection, map[string]vars.Connection, error) {
	mapOfConnections := make(map[string]vars.Connection)
	listOfConnections := make([]vars.Connection, 0)
	_, err := connection.
		NewRequest().
		SetSuccessResult(&mapOfConnections).
		Get("/session/data/{data-source}/connections")

	for _, conn := range mapOfConnections {
		listOfConnections = append(listOfConnections, conn)
	}

	return listOfConnections, mapOfConnections, err
}

func (connection *Connection) Create(name, parentIdentifier, protocol, hostname, port, guacdHost, guacdPort, guacdEncryption string) (vars.Connection, error) {
	newConn := vars.Connection{}
	_, err := connection.
		NewRequest().
		SetSuccessResult(&newConn).
		SetBody(&vars.ConnectionCreate{
			ConnectionBasic: vars.ConnectionBasic{
				Name:             name,
				ParentIdentifier: parentIdentifier,
				Protocol:         protocol,
				Attributes: vars.ConnectionAttributes{
					MaxConnections:        "0",
					MaxConnectionsPerUser: "0",
					GuacdHostname:         guacdHost,
					GuacdPort:             guacdPort,
					GuacdEncryption:       guacdEncryption,
				},
			},
			Parameters: vars.ConnectionParameters{
				Hostname: hostname,
				Port:     port,
			},
		}).
		Post("/session/data/{data-source}/connections")

	return newConn, err
}

func (connection *Connection) Details(identifier string) (vars.Connection, error) {
	conn := vars.Connection{}
	_, err := connection.
		NewRequest().
		SetSuccessResult(&conn).
		SetPathParam("identifier", identifier).
		Get("/session/data/{data-source}/connections/{identifier}")

	return conn, err
}

func (connection *Connection) Update(identifier, name, parentIdentifier, protocol, hostname, port, guacdPort, guacdHost, guacdEncryption string) error {
	_, err := connection.
		NewRequest().
		SetPathParam("identifier", identifier).
		SetBody(&vars.ConnectionCreate{
			ConnectionBasic: vars.ConnectionBasic{
				Name:             name,
				ParentIdentifier: parentIdentifier,
				Protocol:         protocol,
				Attributes: vars.ConnectionAttributes{
					MaxConnections:        "0",
					MaxConnectionsPerUser: "0",
					GuacdHostname:         guacdHost,
					GuacdPort:             guacdPort,
					GuacdEncryption:       guacdEncryption,
				},
			},
			Parameters: vars.ConnectionParameters{
				Hostname: hostname,
				Port:     port,
			},
		}).
		Put("/session/data/{data-source}/connections/{identifier}")

	return err
}

func (connection *Connection) Delete(identifier string) error {
	_, err := connection.
		NewRequest().
		SetPathParam("identifier", identifier).
		Delete("/session/data/{data-source}/connections/{identifier}")

	return err
}
