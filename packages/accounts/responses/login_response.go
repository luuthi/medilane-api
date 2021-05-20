package responses

import (
	"medilane-api/models"
)

type LoginResponse struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	Exp          int64       `json:"exp"`
	User         models.User `json:"user"`
}

func NewLoginResponse(token, refreshToken string, exp int64, user models.User) *LoginResponse {
	return &LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Exp:          exp,
		User:         user,
	}
}
