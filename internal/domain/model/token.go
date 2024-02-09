package model

import "context"

type GetTokenRequest struct {
	Guid string `json:"guid"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRepository interface {
	Create(c context.Context, guid string, refreshToken string) error
	Get(c context.Context, guid string) (string, error)
}
