package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/internal/domain/model"
)

const (
	tokensCollection = "tokens"
)

type refreshTokenRepository struct {
	db         *mongo.Database
	collection string
}

func NewRefreshTokenRepository(db *mongo.Database) model.RefreshTokenRepository {
	return &refreshTokenRepository{
		db: db, collection: tokensCollection,
	}
}

func (r refreshTokenRepository) Create(c context.Context, guid string, refreshToken string) error {
	//TODO implement me
	panic("implement me")
}

func (r refreshTokenRepository) Get(c context.Context, guid string) (string, error) {
	//TODO implement me
	panic("implement me")
}
