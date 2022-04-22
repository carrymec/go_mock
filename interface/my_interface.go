package _interface

import (
	"fmt"
	"github.com/google/uuid"
)

//go:generate mockgen -package=mock -mock_names=Client=MockClient -destination=mock_foo_client.go carrymec/go_mock FooInterface

type FooInterface interface {
	Get(name string) (string, error)
	Set(user FooUser) (FooUser, error)
}

type FooUser struct {
	Id   string
	Name string
	Age  int
}

type Foo struct {
}

func (u *Foo) Set(user FooUser) (FooUser, error) {
	if user.Name == "" {
		return FooUser{}, fmt.Errorf("param name is nil")
	}
	return FooUser{
		Id:   uuid.New().String(),
		Name: user.Name,
		Age:  user.Age,
	}, nil
}

func (u *Foo) Get(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("param name is nil")
	}
	return name, nil
}
