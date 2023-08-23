package client

import "guacamole/vars"

func (client *Client) ConnectionList() ([]vars.ConnectionResp, error) {
	mapOfConnections := make(map[string]vars.ConnectionResp)
	listOfConnections := make([]vars.ConnectionResp, 0)
	_, err := client.
		NewRequest().
		SetSuccessResult(&mapOfConnections).
		Get("/session/data/{data-source}/connections")

	for _, conn := range mapOfConnections {
		listOfConnections = append(listOfConnections, conn)
	}

	return listOfConnections, err
}
