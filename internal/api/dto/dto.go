package dto

type CurrentUser struct {
	Username string
}

func NewCurrentUser(username string) CurrentUser {
	return CurrentUser{
		Username: username,
	}
}
