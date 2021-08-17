package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	users "www.github.com/Ammce/banking-go-auth/domain/Users"
)

type UserRepositoryDB struct {
	db *mongo.Database
}

func (ur UserRepositoryDB) FindBy(email string) (*users.User, *error) {
	var user users.User
	err := ur.db.Collection("users").FindOne(nil, users.User{Email: email}).Decode(user)
	if err != nil {
		return nil, &err
	}
	return &user, nil
}

func (ur UserRepositoryDB) Register(user users.User) (*users.User, *error) {
	res, err := ur.db.Collection("users").InsertOne(nil, &user)

	if err != nil {
		return nil, &err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func NewUserRepositoryDB(db *mongo.Database) *UserRepositoryDB {
	return &UserRepositoryDB{
		db: db,
	}
}
