package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	users "www.github.com/Ammce/banking-go-auth/domain/Users"
)

type UserRepositoryDB struct {
	db *mongo.Database
}

func (ur UserRepositoryDB) Register(user users.User) (*users.User, *error) {
	_, err := ur.db.Collection("users").InsertOne(nil, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})

	if err != nil {
		return nil, &err
	}

	return &user, nil
}

func NewUserRepositoryDB(db *mongo.Database) *UserRepositoryDB {
	return &UserRepositoryDB{
		db: db,
	}
}
