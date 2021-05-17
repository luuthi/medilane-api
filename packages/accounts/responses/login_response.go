package responses

import (
	models2 "medilane-api/packages/accounts/models"
)

type LoginResponse struct {
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
	Exp          int64        `json:"exp"`
	User         models2.User `json:"user"`
}

func NewLoginResponse(token, refreshToken string, exp int64, user models2.User) *LoginResponse {
	return &LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Exp:          exp,
		User:         user,
	}
}
