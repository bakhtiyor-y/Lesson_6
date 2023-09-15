package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/storage"
)

func TestServiceAddUser(t *testing.T) {
	asserts := assert.New(t)

	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login1",
		PasswordHash:     "string",
		Name:             "Name1",
		Surname:          "Surname1",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, err := newService.Add(userData)
	asserts.Nil(err)
	asserts.Equal(1, gotId)
}

func TestServiceAddUser1(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login1",
		PasswordHash:     "string",
		Name:             "Name1",
		Surname:          "Surname1",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	_, err := newService.Add(userData)
	asserts.Nil(err)
	_, err = newService.Add(userData)
	asserts.EqualError(err, "the login/user already exists")
}

func TestServiceGet(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newService.Add(userData)
	got, err := newService.Get(gotId)
	asserts.Nil(err)
	asserts.Equal("Name2", got.Name)
}

func TestServiceGet1(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newService.Add(userData)
	_, err := newService.Get(gotId + 1)
	asserts.EqualError(err, "No user found with id=2")
}

func TestServiceGetUsers1(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login3",
		PasswordHash:     "string",
		Name:             "Name3",
		Surname:          "Surname3",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	_, err := newService.Add(userData)
	if err == nil {
		got, err := newService.GetUsers()
		asserts.Nil(err)
		asserts.Equal("Name3", got[0][2])
	}
}

func TestServiceDelete(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newService.Add(userData)
	err := newService.Delete(gotId)
	asserts.Nil(err)
}

func TestServiceDelete1(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newService.Add(userData)
	err := newService.Delete(gotId + 1)
	asserts.EqualError(err, "No user found with the ID = 2")
}

func TestUpdate(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login4",
		PasswordHash:     "string",
		Name:             "Name4",
		Surname:          "Surname4",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	updated := &models.UserDate{
		Name:       "Name5",
		Surname:    "Surname5",
		Status:     "Active",
		Role:       "User",
		UpdateDate: "string",
	}
	gotId, _ := newService.Add(userData)
	err := newService.Update(gotId, updated)
	asserts.Nil(err)
}

func TestStatistics(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(storage.Storage)
	newStorage.Db = make(map[int]*models.User)
	var newService = New(newStorage)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newService.Add(userData)
	_, err := newService.Get(gotId)
	asserts.Nil(err)
	err = newService.Delete(gotId)
	asserts.Nil(err)

	mapStat := newService.GetStatistics()
	asserts.Equal(1, mapStat["DeletedUsersCount"])
	asserts.Equal(1, mapStat["GetUserCount"])

}
