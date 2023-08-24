package vars

type ConnectionGroupBasic struct {
	Name              string `json:"name"`
	Identifier        string `json:"identifier"`
	Type              string `json:"type"`
	ActiveConnections int    `json:"activeConnections"`
}

type ConnectionGroup struct {
	ConnectionGroupBasic
	ParentIdentifier string        `json:"parentIdentifier"`
	ChildConnections []*Connection `json:"childConnections"`
}

type ConnectionGroupTree struct {
	ConnectionGroupBasic
	ChildGroupConnectionGroups []ConnectionGroup `json:"childConnectionGroups"`
}
