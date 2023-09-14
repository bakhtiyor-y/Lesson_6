package api

import (
	"github.com/Shemistan/Lesson_5/internal/converters"
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/service"
)

type IApi interface {
	Add(req *models.AddRequest) (int, error)
	Get(userId int) (*models.User, error)
	GetUsers() ([][]interface{}, error)
	Update(userId int, user *models.UserDate) error
	Delete(userId int) error
	GetStatistics() map[string]int
}

func New(serv service.IService) *Api {
	return &Api{serv: serv}
}

type Api struct {
	serv service.IService
}

func (a *Api) Add(req *models.AddRequest) (int, error) {
	res, err := a.serv.Add(converters.ApiModelToServiceModel(*req))
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (a *Api) GetUsers() ([][]interface{}, error) {
	res, err := a.serv.GetUsers()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (a *Api) Get(userId int) (*models.User, error) {
	res, err := a.serv.Get(userId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (a *Api) Update(userId int, user *models.UserDate) error {
	err := a.serv.Update(userId, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) Delete(userId int) error {
	err := a.serv.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) GetStatistics() map[string]int {
	return a.serv.GetStatistics()
}
