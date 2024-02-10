package model

import "context"

type GetTokenRequest struct {
	Guid string `json:"guid"`
}

type GetTokensResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type RefreshTokenRepository interface {
	Create(c context.Context, guid string, refreshToken string) error
	Get(c context.Context, guid string) (string, error)
}
