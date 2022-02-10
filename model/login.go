package model

type Login struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}
