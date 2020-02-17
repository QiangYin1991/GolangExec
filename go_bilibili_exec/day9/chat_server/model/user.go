package model

const (
	UserStatusOffline = iota
	UserStatusOnline
)

type User struct {
	UserId int `json:"user_Id"`
	Passwd string `json:"passwod"`
	Nick string `json:"nick"`
	Sex string `json:"sex"`
	Header string `json:"header"`
	LastLogin string `json:"last_login"`
	Status int `json:"status"`
}
