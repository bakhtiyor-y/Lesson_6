package models

type User struct {
	Login            string
	PasswordHash     string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate string
	UpdateDate       string
}
