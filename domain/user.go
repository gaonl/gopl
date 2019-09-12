package domain

type User struct {
	Id               int64
	Name             string
	NickNames        [2]string
	Password         string
	RegisterDateTime int64
}
