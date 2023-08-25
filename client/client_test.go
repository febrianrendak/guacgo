package client_test

import (
	"fmt"
	"github.com/febrianrendak/guacgo/client"
	"github.com/google/uuid"
	"testing"
)

func PrepareClient() *client.Client {
	return client.NewClient("http://192.168.210.171:8080", "guacadmin", "guacadmin")
}

func TestGetToken(t *testing.T) {
	newClient := PrepareClient()

	_, err := newClient.Auth().Token()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetConnectionList(t *testing.T) {
	newClient := PrepareClient()

	connList, connMap, err := newClient.Connection().List()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(connList, connMap)
}

func TestCreateDeleteConnection(t *testing.T) {
	newClient := PrepareClient()

	connName := fmt.Sprintf("conn-%s", uuid.New().String())
	newConn, err := newClient.Connection().Create(
		connName,
		"1",
		"rdp",
		"192.168.210.171",
		"15000",
		"192.168.210.171",
		"4822",
		"none",
	)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(newConn)

	if err := newClient.Connection().Delete(newConn.Identifier); err != nil {
		t.Fatal(err)
	}
}

func TestConnectionGroupTree(t *testing.T) {
	newClient := PrepareClient()

	cgTree, err := newClient.ConnectionGroup().Tree("ROOT")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cgTree)
}

func TestConnectionGroupDetails(t *testing.T) {
	newClient := PrepareClient()

	cg, err := newClient.ConnectionGroup().Details("ROOT")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cg)
}

func TestConnectionGroupList(t *testing.T) {
	newClient := PrepareClient()

	cgList, cgMap, err := newClient.ConnectionGroup().List()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cgList)
	fmt.Println(cgMap)
}

func TestUserLIst(t *testing.T) {
	newClient := PrepareClient()

	userList, mapOfUserList, err := newClient.User().List()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(userList, mapOfUserList)
}

func TestUserDetails(t *testing.T) {
	newClient := PrepareClient()

	user, err := newClient.User().Details("febrian@paques.id")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(user)
}

func TestCreateAndUpdateUser(t *testing.T) {
	var err error
	var newClient = PrepareClient()

	user, err := newClient.User().Create(
		"febrian2@paques.id",
		"febrian2@gmail.com",
		"Febrian Rendak",
		"true",
		"",
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(user)

	err = newClient.User().Update(
		"febrian2@paques.id",
		"febrian2@gmail.com",
		"Febrian Rendak",
		"false",
		"paques123",
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserGroupList(t *testing.T) {
	newClient := PrepareClient()

	userGroupList, mapOfUserGroups, err := newClient.UserGroup().List()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(userGroupList, mapOfUserGroups)
}

func TestGetUserGroupDetails(t *testing.T) {
	newClient := PrepareClient()

	userGroup, err := newClient.
		UserGroup().
		Details("MAIN")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(userGroup)
}

func TestGetUserGroupPermission(t *testing.T) {
	newClient := PrepareClient()

	permissions, err := newClient.
		UserGroup().
		Permissions("MAIN")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(permissions)
}

func TestUserGroupOperation(t *testing.T) {
	newClient := PrepareClient()

	err := newClient.
		UserGroup().
		ConnectionOperation("MAIN", "1", "add")

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserUserGroupOperation(t *testing.T) {
	newClient := PrepareClient()

	err := newClient.
		User().
		UserGroupOperation("febrian2@paques.id", "MAIN", "add")

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetConnectionParameters(t *testing.T) {
	newClient := PrepareClient()

	connAttributes, err := newClient.
		Connection().
		Parameters("1")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(connAttributes)
}
