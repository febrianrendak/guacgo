package vars

type ConnectionParametersMap map[string]string

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
	Parameters ConnectionParametersMap `json:"parameters"`
}

type Connection struct {
	ConnectionBasic
	Identifier        string                  `json:"identifier"`
	ActiveConnections int                     `json:"activeConnections"`
	LastActive        int                     `json:"lastActive"`
	Parameters        ConnectionParametersMap `json:"parameters,omitempty"`
}
