package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/internal/domain/model"
)

const (
	tokensCollection = "tokens"
)

type RefreshTokenRepository struct {
	db         *mongo.Database
	collection string
}

func NewRefreshTokenRepository(db *mongo.Database) model.RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db, collection: tokensCollection,
	}
}

func (r RefreshTokenRepository) Create(c context.Context, guid string, refreshToken string) error {
	collection := r.db.Collection(tokensCollection)
	token := model.TokenEntity{
		GUID:             guid,
		RefreshTokenHash: refreshToken,
	}

	_, err := collection.InsertOne(c, token)
	if err != nil {
		return err
	}

	return nil
}

func (r RefreshTokenRepository) GetByGuid(c context.Context, guid string) (*model.TokenEntity, error) {
	collection := r.db.Collection(tokensCollection)

	var token model.TokenEntity

	err := collection.FindOne(c, primitive.M{"guid": guid}).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
