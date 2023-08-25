package vars

/*
{
    "identifier": "4c78c580-88e8-3a2d-b47e-34863e7b5f07",
    "connectionIdentifier": "1",
    "startDate": 1692934377093,
    "remoteHost": "192.168.210.71",
    "username": "guacadmin",
    "connectable": true
}
*/

type ActiveConnection struct {
	Identifier           string `json:"identifier"`
	ConnectionIdentifier string `json:"connectionIdentifier"`
	StartDate            int    `json:"startDate"`
	RemoteHost           string `json:"remoteHost"`
	Username             string `json:"username"`
	Connectable          bool   `json:"connectable"`
}
