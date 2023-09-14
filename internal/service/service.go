package service

import (
	"database/sql"
	"fmt"

	"github.com/Shemistan/Lesson_5/internal/converters"
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/storage"
)

type IService interface {
	Add(user *models.User) (int, error)
	Get(userId int) (*models.User, error)
	GetUsers() ([][]interface{}, error)
	Update(userId int, user *models.UserDate) error
	Delete(userId int) error
	GetStatistics() map[string]int
}

func New(repo storage.IStorage) *Service {
	return &Service{
		CounterAdd:      0,
		CounterGet:      0,
		CounterGetUsers: 0,
		CounterUpdate:   0,
		CounterDelete:   0,
		repo:            repo,
	}
}

type Service struct {
	CounterAdd      int
	CounterGet      int
	CounterGetUsers int
	CounterUpdate   int
	CounterDelete   int
	repo            storage.IStorage
}

func (s *Service) Add(user *models.User) (int, error) {
	res, err := s.repo.Add(user)
	if err == sql.ErrNoRows {
		fmt.Println("no rows")
	}

	if err != nil {
		return 0, err
	}

	s.CounterAdd++

	return res, nil
}

func (s *Service) Get(userId int) (*models.User, error) {
	res, err := s.repo.Get(userId)
	if err != nil {
		return res, err
	}

	s.CounterGet++

	return res, nil
}

func (s *Service) Update(userId int, user *models.UserDate) error {
	err := s.repo.Update(userId, user)
	if err != nil {
		return err
	}

	s.CounterUpdate++

	return nil
}

func (s *Service) Delete(userId int) error {
	err := s.repo.Delete(userId)
	if err != nil {
		return err
	}

	s.CounterDelete++

	return nil
}

func (s *Service) GetUsers() ([][]interface{}, error) {
	res, _ := s.repo.GetUsers()

	result := converters.MapToSlice(res)

	s.CounterGetUsers++

	return result, nil
}

func (s *Service) GetStatistics() map[string]int {
	statMap := make(map[string]int)
	statMap["GetUserCount"] = s.CounterGet
	statMap["GetUsersCount"] = s.CounterGetUsers
	statMap["UpdateCount"] = s.CounterUpdate
	statMap["DeletedUsersCount"] = s.CounterDelete

	return statMap
}
