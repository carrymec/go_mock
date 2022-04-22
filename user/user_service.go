package user

import (
	"fmt"

	mytest "carrymec/go_mock/interface"
)

type Service struct {
	client mytest.FooInterface
}

func (s Service) SetUser(user mytest.FooUser) (string, error) {
	fooUser, err := s.client.Set(user)
	if err != nil {
		return "", err
	}
	return fooUser.Id, nil
}

func (s Service) GetUser(name string) (string, error) {
	user, err := s.client.Get(name)
	if err != nil {
		return "", err
	}
	if user == "" {
		user, err = s.client.Get(name)
		if err != nil {
			return "", err
		}
	}
	return user, nil
}

func Clear() {
	fmt.Println("ok")
}

// 获取用户信息通过用户名 并且返回该用户信息
func (s Service) GetUserByName(name string) (string, error) {
	user, err := s.client.Get(name)
	if err != nil {
		return "", err
	}
	if user == "" {
		user, err = s.client.Get(name)
		if err != nil {
			return "", err
		}
	}
	return user, nil
}
