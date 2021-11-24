package domain

type Object struct {
	Token string `json:"access_token"`
	Scope string `json:"scope"`
	Exp string `json:"exp"`
	TokenType string `json:"token_type"`
}
