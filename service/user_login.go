package service

import "panLite/resp"

type LoginService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (service *LoginService) PlainTextLogin() *resp.Error {
	return resp.ReqInvalidErr
}
