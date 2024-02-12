package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r RefreshTokenRepository) UpsertToken(c context.Context, guid string, refreshToken string) error {
	collection := r.db.Collection(tokensCollection)

	filter := bson.M{"guid": guid}
	update := bson.M{
		"$set": bson.M{
			"refreshTokenHash": refreshToken,
		},
	}
	option := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(c, filter, update, option)
	if err != nil {
		return err
	}

	return nil
}
