package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	GetByGuid(c context.Context, guid string) (*TokenEntity, error)
}

type TokenEntity struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	GUID             string             `bson:"guid"`
	RefreshTokenHash string             `bson:"refreshTokenHash"`
}
