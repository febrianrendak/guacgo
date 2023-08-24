package vars

type UserGroupBasic struct {
	Identifier string `json:"identifier"`
}

type UserGroupAttributes struct {
	Disabled string `json:"disabled"`
}

type UserGroup struct {
	UserGroupBasic
	Attributes UserGroupAttributes `json:"attributes"`
}

type UserGroupPermission struct {
	ConnectionPermissions      map[string][]string `json:"connectionPermissions"`
	ConnectionGroupPermissions map[string][]string `json:"connectionGroupPermissions"`
}
