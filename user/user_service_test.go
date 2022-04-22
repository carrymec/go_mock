package user

import (
	_interface "carrymec/go_mock/interface"
	"carrymec/go_mock/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestService_SetUser(t *testing.T) {
	type args struct {
		user _interface.FooUser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test failed",
			args: args{user: _interface.FooUser{
				Name: "chen",
				Age:  22,
			}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			fooInterface := mocks.NewMockFooInterface(ctrl)
			fooInterface.EXPECT().Set(tt.args.user).Times(1).Return(_interface.FooUser{
				Id:   uuid.New().String(),
				Name: "chen",
				Age:  22,
			}, fmt.Errorf("400"))
			s := Service{
				client: fooInterface,
			}
			got, err := s.SetUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUser(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{name: "chen"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			fooInterface := mocks.NewMockFooInterface(ctrl)
			gomock.InOrder(
				fooInterface.EXPECT().Get(tt.args.name).Times(1).Return("", nil),
				fooInterface.EXPECT().Get(tt.args.name).Times(1).Return("", errors.New("500")),
			)
			s := Service{
				client: fooInterface,
			}
			got, err := s.GetUser(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}


