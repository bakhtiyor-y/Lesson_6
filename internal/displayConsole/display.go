package displayConsole

import (
	"fmt"
	"time"

	"github.com/Shemistan/Lesson_5/internal/models"
)

func RegistrationDataFromConsole() (*models.User, error) {
	var initialData models.User

	fmt.Println("Please enter login")
	_, err := fmt.Scan(&initialData.Login)
	if err != nil {
		return nil, err
	}

	fmt.Println("Please enter password")
	_, err = fmt.Scan(&initialData.PasswordHash)
	if err != nil {
		return nil, err
	}

	fmt.Println("Please enter name")
	_, err = fmt.Scan(&initialData.Name)
	if err != nil {
		return nil, err
	}

	fmt.Println("Please enter surname")
	_, err = fmt.Scan(&initialData.Surname)
	if err != nil {
		return nil, err
	}

	return &initialData, nil
}

func GetUserIdFromConsole() (int, error) {
	var userId int

	fmt.Println("Please enter user ID")
	_, err := fmt.Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func GetDataFromConsoleToUpdate() (int, *models.UserDate, error) {
	var userId int
	var fromInput string
	var userData models.UserDate

	fmt.Println("Please, enter user ID")
	_, err := fmt.Scan(&userId)
	if err != nil {
		return 0, nil, err
	}

	fmt.Println("Please enter name")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		return 0, nil, err
	}
	userData.Name = fromInput

	fmt.Println("Please enter surname")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		return 0, nil, err
	}
	userData.Surname = fromInput

	fmt.Println("Please enter status")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		return 0, nil, err
	}
	userData.Status = fromInput

	fmt.Println("Please enter role")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		return 0, nil, err
	}
	userData.Role = fromInput

	userData.UpdateDate = time.Now().Format("01-02-2006 15:04:05")

	return userId, &userData, nil
}

func GetDataFromConsoleToDelete() int {
	var userId int

	fmt.Println("Please, enter user ID")
	_, err := fmt.Scan(&userId)
	if err != nil {
		fmt.Println(err)
	}

	return userId
}

func Line() {
	fmt.Println("=============")
}
