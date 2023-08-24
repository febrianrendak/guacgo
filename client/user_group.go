package client

import (
	"fmt"
	"github.com/febrianrendak/guacgo/vars"
)

type UserGroup struct {
	*Client
}

func (client *Client) UserGroup() *UserGroup {
	return &UserGroup{
		client,
	}
}

func (ug *UserGroup) List() (userGroupList []vars.UserGroup, mapOfUserGroups map[string]vars.UserGroup, err error) {
	userGroupList = make([]vars.UserGroup, 0)
	mapOfUserGroups = make(map[string]vars.UserGroup)

	_, err = ug.NewRequest().
		SetSuccessResult(&mapOfUserGroups).
		Get("/guacamole/api/session/data/{data-source}/userGroups")

	for _, userGroup := range mapOfUserGroups {
		userGroupList = append(userGroupList, userGroup)
	}

	return
}

func (ug *UserGroup) Details(identifier string) (userGroup vars.UserGroup, err error) {
	_, err = ug.NewRequest().
		SetSuccessResult(&userGroup).
		SetPathParam("identifier", identifier).
		Get("/guacamole/api/session/data/{data-source}/userGroups/{identifier}")

	return
}

func (ug *UserGroup) Permissions(identifier string) (userGroupPermissions vars.UserGroupPermission, err error) {
	_, err = ug.NewRequest().
		SetSuccessResult(&userGroupPermissions).
		SetPathParam("identifier", identifier).
		Get("/guacamole/api/session/data/{data-source}/userGroups/{identifier}/permissions")

	return
}

// ConnectionOperation add or remove connection from an user group
func (ug *UserGroup) ConnectionOperation(identifier, connectionIdentifier, op string) (err error) {
	resp, err := ug.NewRequest().
		SetPathParam("identifier", identifier).
		SetBody([]vars.Operation{
			{
				OP:    op,
				Path:  fmt.Sprintf("/connectionPermissions/%s", connectionIdentifier),
				Value: "READ",
			},
		}).
		Patch("/guacamole/api/session/data/{data-source}/userGroups/{identifier}/permissions")

	fmt.Println(resp)
	return
}
