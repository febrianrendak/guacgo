package vars

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

type ConnectionBasic struct {
	Name             string               `json:"name"`
	ParentIdentifier string               `json:"parentIdentifier"`
	Protocol         string               `json:"protocol"`
	Attributes       ConnectionAttributes `json:"attributes"`
}

type ConnectionCreate struct {
	ConnectionBasic
	Parameters ConnectionParameters `json:"parameters"`
}

type Connection struct {
	ConnectionBasic
	Identifier        string               `json:"identifier"`
	ActiveConnections int                  `json:"activeConnections"`
	LastActive        int                  `json:"lastActive"`
	Parameters        ConnectionParameters `json:"parameters,omitempty"`
}
