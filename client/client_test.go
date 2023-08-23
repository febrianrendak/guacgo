package client_test

import (
	"fmt"
	"guacamole/client"
	"testing"
)

func TestGetToken(t *testing.T) {
	newClient := client.NewClient("http://192.168.210.171:8080/guacamole/api", "guacadmin", "guacadmin")

	_, err := newClient.Auth().Token()
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetConnectionList(t *testing.T) {
	newClient := client.NewClient("http://192.168.210.171:8080/guacamole/api", "guacadmin", "guacadmin")

	connList, connMap, err := newClient.Connection().List()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(connList, connMap)
}
