package client

import "github.com/febrianrendak/guacgo/vars"

type User struct {
	*Client
}

func (client *Client) User() *User {
	return &User{
		client,
	}
}

func (u *User) List() (userList []vars.User, mapOfUsers map[string]vars.User, err error) {
	userList = make([]vars.User, 0)
	mapOfUsers = make(map[string]vars.User)

	_, err = u.NewRequest().
		SetSuccessResult(&mapOfUsers).
		Get("/api/session/data/{data-source}/users")

	for _, user := range mapOfUsers {
		userList = append(userList, user)
	}

	return
}

func (u *User) Details(username string) (user vars.User, err error) {
	_, err = u.NewRequest().
		SetSuccessResult(&user).
		SetPathParam("username", username).
		Get("/api/session/data/{data-source}/users/{username}")

	return
}

func (u *User) Create(username, guacEmailAddress, guacFullName, disabled, password string) (user vars.User, err error) {
	_, err = u.NewRequest().
		SetSuccessResult(&user).
		SetBody(&vars.UserCreate{
			UserBasic: vars.UserBasic{
				Username: username,
			},
			Password: password,
			Attributes: vars.UserAttributes{
				GuacEmailAddress: guacEmailAddress,
				GuacFullName:     guacFullName,
				Disabled:         disabled,
			},
		}).
		Post("/api/session/data/{data-source}/users")

	return
}

func (u *User) Update(username, newUsername, guacEmailAddress, guacFullName, disabled, password string) (err error) {
	_, err = u.NewRequest().
		SetPathParam("username", username).
		SetBody(&vars.UserCreate{
			UserBasic: vars.UserBasic{
				Username: newUsername,
			},
			Password: password,
			Attributes: vars.UserAttributes{
				GuacEmailAddress: guacEmailAddress,
				GuacFullName:     guacFullName,
				Disabled:         disabled,
			},
		}).
		Put("/api/session/data/{data-source}/users/{username}")

	if err == nil && username != "" && newUsername != "" {
		if err = UpdateUsername(username, newUsername); err != nil {
			return
		}
	}

	return
}

// UserGroupOperation add or remove user from an group where OP either add or remove
func (u *User) UserGroupOperation(username, userGroup, op string) (err error) {
	_, err = u.NewRequest().
		SetPathParam("username", username).
		SetBody([]vars.Operation{
			{
				OP:    op,
				Path:  "/",
				Value: userGroup,
			},
		}).
		Patch("/api/session/data/{data-source}/users/{username}/userGroups")
	return
}
