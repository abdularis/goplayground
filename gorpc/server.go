package main

import (
	"context"
	"errors"
	"gorpc/service"
)


var usersDB = map[string]service.UserDetail {
	"usr1": {
		UserID:   "usr1",
		Name:     "aris",
		Email:    "aris@mail.com",
		ImageUrl: "",
	},
	"usr2": {
		UserID:   "usr2",
		Name:     "abdul",
		Email:    "abdul@mail.com",
		ImageUrl: "",
	},
}


type UserServiceServerImpl struct {
}

func (s *UserServiceServerImpl) GetUserDetail(ctx context.Context, req *service.UserDetailRequest) (*service.UserDetail, error) {
	if userDetail, ok := usersDB[req.UserID]; ok {
		return &userDetail, nil
	}
	return nil, errors.New("user not found")
}

