package vars

type UserBasic struct {
	Username string `json:"username"`
}

type UserAttributes struct {
	GuacEmailAddress string `json:"guac-email-address,omitempty"`
	GuacFullName     string `json:"guac-full-name,omitempty"`
	Disabled         string `json:"disabled,omitempty"`
}

type User struct {
	UserBasic
	Attributes UserAttributes `json:"attributes"`
	LastActive int            `json:"lastActive"`
}

type UserCreate struct {
	UserBasic
	Password   string         `json:"password,omitempty"`
	Attributes UserAttributes `json:"attributes"`
}

type UserDisabled struct {
	UserBasic
	Attributes UserAttributes `json:"attributes"`
}
