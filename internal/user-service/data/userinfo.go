package data

type UserInfo []map[string]interface{}

type UserInfoInterface interface {
	GetUserInfo(userid string) (UserInfo, error)
	SetUserInfo(user UserInfo) error
	AddUserInfo(user UserInfo) (string, error)
}
