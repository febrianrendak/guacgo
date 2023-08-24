package client_test

import (
	"fmt"
	"github.com/google/uuid"
	"guacamole/client"
	"testing"
)

func PrepareClient() *client.Client {
	return client.NewClient("http://192.168.210.171:8080/guacamole/api", "guacadmin", "guacadmin")
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
