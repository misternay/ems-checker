package model

type GetToken struct {
	Expire string `json:"expire"`
	Token  string `json:"token"`
}
