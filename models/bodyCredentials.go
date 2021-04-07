package models

//swagger:model body_credentials
type BodyCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

