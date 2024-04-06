package storage

import (
	"taskmaster/entity"
	"testing"
)

var UUser_id uint
var UToken []string

func TestUserCreate(t *testing.T) {
	newUser := entity.User{
		Login:    "testuser",
		Email:    "testuser@example.com",
		Role:     "user",
		Password: "testpassword",
	}

	createdUser := UserCreate(newUser)
	UUser_id = createdUser.ID
	if createdUser == nil {
		t.Error("User create failed")
	}

	if createdUser.Login != newUser.Login {
		t.Errorf("Expected login %s, got %s", newUser.Login, createdUser.Login)
	}
}

func TestUserGetAll(t *testing.T) {
	users := UserGetAll()

	if users == nil {
		t.Error("User get all failed")
	}
}

func TestUserGet(t *testing.T) {
	user := UserGet(UUser_id)

	if user == nil {
		t.Errorf("Failed to get user with ID %d", UUser_id)
	}
}

func TestUserUpdate(t *testing.T) {
	updateUser := entity.User{
		Login:    "updateduser",
		Email:    "updateduser@example.com",
		Role:     "user",
		Password: "updatedpassword",
	}

	updatedUser := UserUpdate(updateUser, UUser_id)

	if updatedUser == nil {
		t.Errorf("Failed to update user with ID %d", UUser_id)
	}
}

func TestUserAuth(t *testing.T) {
	newUser := entity.User{
		Login:    "testuser",
		Email:    "testuser@example.com",
		Role:     "user",
		Password: "testpassword",
	}

	authUserToken := UserAuth(newUser)
	UToken = []string{authUserToken.Token}
	if authUserToken == nil {
		t.Error("User authentication failed")
	}
}

func TestUserAuthDelete(t *testing.T) {
	deletedToken := UserAuthDelete(UToken)

	if deletedToken == nil {
		t.Errorf("Failed to delete token %s", UToken)
	}
}

func TestUserDelete(t *testing.T) {
	deletedUser := UserDelete(UUser_id)

	if deletedUser == nil {
		t.Errorf("Failed to delete user with ID %d", UUser_id)
	}
}
