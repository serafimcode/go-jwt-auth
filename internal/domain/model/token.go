package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTokenRequest struct {
	Guid string `json:"guid"`
}

type RefreshTokenRequest struct {
	AccessToken string `json:"accessToken"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type TokensPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRepository interface {
	UpsertToken(c context.Context, guid string, refreshToken string) error
}

type TokenEntity struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	GUID             string             `bson:"guid"`
	RefreshTokenHash string             `bson:"refreshTokenHash"`
}
