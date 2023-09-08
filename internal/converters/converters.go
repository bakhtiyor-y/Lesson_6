package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Shemistan/Lesson_5/internal/models"
)

func ApiModelToServiceModel(req models.AddRequest) *models.User {
	hash := sha256.New()
	hash.Write([]byte(req.AuthParams.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)
	res := &models.User{
		Login:            req.AuthParams.Login,
		PasswordHash:     hashPassword,
		Name:             req.Date.Name,
		Surname:          req.Date.Surname,
		Status:           "active",
		Role:             "User",
		RegistrationDate: req.Date.RegistrationDate,
		UpdateDate:       req.Date.UpdateDate,
	}

	return res
}

func MapToSlice(initialMap map[int]*models.User) [][]interface{} {
	sliceOfSlices := make([][]interface{}, 0)
	//
	for k, v := range initialMap {
		slice := []interface{}{k, v.Login, v.Name, v.Surname, v.Status, v.Role, v.RegistrationDate, v.UpdateDate}
		sliceOfSlices = append(sliceOfSlices, slice)
	}
	return sliceOfSlices
}
