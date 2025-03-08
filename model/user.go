package model

type UserAuth struct {
	UserName string
	Password string
}

type UserType uint8

const (
	Admin UserType = iota + 1
	Normal
)
