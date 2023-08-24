package client

import "guacamole/vars"

type ConnectionGroup struct {
	*Client
}

func (client *Client) ConnectionGroup() *ConnectionGroup {
	return &ConnectionGroup{
		client,
	}
}

func (cg *ConnectionGroup) List() (connGroupList []vars.ConnectionGroup, mapOfConnGroup map[string]vars.ConnectionGroup, err error) {
	connGroupList = make([]vars.ConnectionGroup, 0)
	mapOfConnGroup = make(map[string]vars.ConnectionGroup)

	_, err = cg.NewRequest().
		SetSuccessResult(&mapOfConnGroup).
		Get("/session/data/{data-source}/connectionGroups")

	for _, connGroup := range mapOfConnGroup {
		connGroupList = append(connGroupList, connGroup)
	}

	return
}

func (cg *ConnectionGroup) Tree(identifier string) (connGroupTree vars.ConnectionGroupTree, err error) {
	_, err = cg.NewRequest().
		SetSuccessResult(&connGroupTree).
		SetPathParam("identifier", identifier).
		Get("/session/data/{data-source}/connectionGroups/{identifier}/tree")

	return
}

func (cg *ConnectionGroup) Details(identifier string) (connGroup vars.ConnectionGroup, err error) {
	_, err = cg.NewRequest().
		SetSuccessResult(&connGroup).
		SetPathParam("identifier", identifier).
		Get("/session/data/{data-source}/connectionGroups/{identifier}")

	return
}
