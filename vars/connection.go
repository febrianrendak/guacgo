package vars

/*
"guacd-encryption": "none",
"max-connections": "1",
"guacd-hostname": "192.168.210.171",
"guacd-port": "4822",
"max-connections-per-user": "1"
*/

type ConnectionParameters struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

type ConnectionAttributes struct {
	MaxConnections        string `json:"max-connections"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
	GuacdHostname         string `json:"guacd-hostname"`
	GuacdPort             string `json:"guacd-port"`
	GuacdEncryption       string `json:"guacd-encryption"`
}

type Connection struct {
	Name             string               `json:"name"`
	ParentIdentifier string               `json:"parentIdentifier"`
	Protocol         string               `json:"protocol"`
	Attributes       ConnectionAttributes `json:"attributes"`
}

type ConnectionReq struct {
	Connection
	Parameters ConnectionParameters `json:"parameters"`
}

type ConnectionResp struct {
	Identifier string `json:"identifier"`
	Connection
}
