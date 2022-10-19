package User_test

import (
	"mock/User"
	mock "mock/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUser := mock.NewMockMako(mockCtrl)
	testUser := &User.User{IUser: mockUser}

	mockUser.EXPECT().AddUser(1, "sample mock test").Return(nil).Times(1)
	testUser.Use()

}
