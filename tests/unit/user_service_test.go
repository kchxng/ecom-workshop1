package unit_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	app_user "github.com/kchxng/ecom-api/api/users"
	"github.com/kchxng/ecom-api/configs/db"
	"github.com/kchxng/ecom-api/models"
)

func TestCreateUser(t *testing.T) {
	db := db.InitDatabase()
	repo := app_user.NewUserRepository(db)
	svc := app_user.NewUserService(repo)
	var data models.Users
	_, err := svc.Create(&data)
	if err != nil {
		t.Error(err)
	}

	// assert.Equal(t, 1, user.Id)
	// assert.Equal(t, "Alice", user.Name)
}

func TestGetUser(t *testing.T) {
	db := db.InitDatabase()
	repo := app_user.NewUserRepository(db)
	svc := app_user.NewUserService(repo)
	fetchedUser, err := svc.GetSingle("1")
	t.Error(t, err)
	assert.Equal(t, "Bob", fetchedUser.Id)
}

func TestUpdateUser(t *testing.T) {
	db := db.InitDatabase()
	repo := app_user.NewUserRepository(db)
	svc := app_user.NewUserService(repo)
	err := svc.Update("1", nil)
	t.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	db := db.InitDatabase()
	repo := app_user.NewUserRepository(db)
	svc := app_user.NewUserService(repo)
	err := svc.Delete("1")
	t.Error(t, err)
}
