package storage

import "github.com/Devon-ODell/testcode/models"

type MongoStorage struct{}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID: 1,
		Name: "Foo",
	}
}



