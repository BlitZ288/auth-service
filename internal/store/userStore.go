package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoUserStore struct {
	col *mongo.Collection
}

func (store *mongoUserStore) Create(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now().UTC()
	_, err := store.col.InsertOne(ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		return ErrDuplicateEmail
	}
	return err
}

func (store *mongoUserStore) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := store.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNotFound
	}
	return &user, err
}

func (store *mongoUserStore) FindById(ctx context.Context, id primitive.ObjectID) (*User, error) {
	var user User
	err := store.col.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNotFound
	}
	return &user, err
}

func (store *mongoUserStore) Update(ctx context.Context, user *User, update bson.M) error {
	result, err := store.col.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func (store *mongoUserStore) EnsureIndexes(ctx context.Context) error {
	model := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := store.col.Indexes().CreateOne(ctx, model)
	if err != nil {
		return fmt.Errorf("create indexes: %w", err)
	}
	return nil
}
