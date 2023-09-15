package storage

import (
	"errors"
	"strconv"

	"github.com/Shemistan/Lesson_5/internal/models"
)

type IStorage interface {
	Add(user *models.User) (int, error)
	Get(userId int) (*models.User, error)
	GetUsers() (map[int]*models.User, error)
	Update(userId int, user *models.UserDate) error
	Delete(userID int) error
}

type Storage struct {
	Db   map[int]*models.User
	ids  int
	Host string
	Port int
	TLL  int
	conn *Conn
}

func New(host string, port, ttl int, conn *Conn) IStorage {
	return &Storage{
		Db:   make(map[int]*models.User),
		ids:  0,
		Host: host,
		Port: port,
		TLL:  ttl,
		conn: conn,
	}
}

func (s *Storage) Add(user *models.User) (int, error) {
	if user == nil {
		return 0, errors.New("user is nil")
	}

	s.ids++

	for k, v := range s.Db {
		if v.Login == user.Login {
			return k, errors.New("the login/user already exists")
		}
	}

	s.Db[s.ids] = user

	return s.ids, nil
}

func (s *Storage) Get(userId int) (*models.User, error) {
	user, ok := s.Db[userId]
	if !ok {
		err := errors.New("No user found with id=" + strconv.Itoa(userId))
		return user, err
	}

	res := &models.User{
		Login:            user.Login,
		Name:             user.Name,
		Surname:          user.Surname,
		Status:           user.Status,
		Role:             user.Role,
		RegistrationDate: user.RegistrationDate,
		UpdateDate:       user.UpdateDate,
	}

	return res, nil
}

func (s *Storage) GetUsers() (map[int]*models.User, error) {
	return s.Db, nil
}

func (s *Storage) Update(userId int, user *models.UserDate) error {
	userInDb, ok := s.Db[userId]
	if !ok {
		err := errors.New("No user found with id=" + strconv.Itoa(userId))
		return err
	}

	userInDb.Name = user.Name
	userInDb.Surname = user.Surname
	userInDb.Status = user.Status
	userInDb.Role = user.Role
	userInDb.UpdateDate = user.UpdateDate

	return nil
}

func (s *Storage) Delete(userID int) error {
	_, ok := s.Db[userID]
	if ok {
		delete(s.Db, userID)
		return nil
	}

	err := errors.New("No user found with the ID = " + strconv.Itoa(userID))
	return err
}

func NewConnect() *Conn {
	return &Conn{}
}

type Conn struct {
	val bool
}

func (c *Conn) IsClose() bool {
	return !(c.val)
}
