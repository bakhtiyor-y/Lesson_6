package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shemistan/Lesson_5/internal/api"
	"github.com/Shemistan/Lesson_5/internal/displayConsole"
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/service"
	"github.com/Shemistan/Lesson_5/internal/storage"
	"os"
	"strconv"
	"time"
)

func main() {

	conn := storage.NewConnect()
	db := storage.New("localhost", 5432, 30, conn)
	serv := service.New(db)
	handleRes := api.New(serv)
	Server(handleRes)
}

func Server(handleRes api.IApi) {
	var choice int
	for {
		fmt.Println("\nSelect the option (enter the number): \n1.Register user. 2.Get user. 3.Update user. 4.Delete user. 5.Get all users. 6.GetStatistics. 7.Exit")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
		}
		switch choice {
		case 1:
			initialData := displayConsole.RegistrationDataFromConsole()
			addUser(handleRes, initialData)
		case 2:
			fmt.Println("case2 Get user")
			UserIdFromConsole := displayConsole.GetUserIdFromConsole()
			getUser(handleRes, UserIdFromConsole)
		case 3:
			fmt.Println("case3 update")
			userId, userData := displayConsole.GetDataFromConsoleToUpdate()
			updateUser(handleRes, userId, userData)
		case 4:
			fmt.Println("case4 Delete")
			userId := displayConsole.GetDataFromConsoleToDelete()
			deleteUser(handleRes, userId)
		case 5:
			fmt.Println("case4 GetAllUsers")
			//slice , err := api.Api.GetUsers()
			getUsers(handleRes)
		case 6:
			getStat(handleRes)
		case 7:
			os.Exit(0)
		default:
			fmt.Println("Enter 1-7")
		}
	}
}

func addUser(handleRes api.IApi, initialData models.User) {
	res, err := handleRes.Add(&models.AddRequest{
		AuthParams: models.UserAuthParams{
			Login:    initialData.Login,
			Password: initialData.PasswordHash,
		},
		Date: models.UserDate{
			Name:             initialData.Name,
			Surname:          initialData.Surname,
			RegistrationDate: time.Now().Format("01-02-2006 15:04:05"),
		},
	})
	if err != nil {
		displayConsole.Line()
		fmt.Println(err)
		displayConsole.Line()
	} else {
		displayConsole.Line()
		fmt.Println("\nUser has been registered with id = " + strconv.Itoa(res) + ";")
		displayConsole.Line()
	}

}

func getUser(handleRes api.IApi, userId int) {
	res, err := handleRes.Get(userId)
	if err != nil {
		fmt.Println(err)
	}

	jsonStr, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		displayConsole.Line()
		fmt.Println(string(jsonStr))
		displayConsole.Line()
	}
}

func updateUser(handleRes api.IApi, userId int, userDate *models.UserDate) {
	err := handleRes.Update(userId, userDate)
	if err != nil {
		fmt.Println(err)
	}
	displayConsole.Line()
	fmt.Printf("The user with ID= %v was successfully updated\n", userId)
	displayConsole.Line()
}

func deleteUser(handleRes api.IApi, userId int) {
	err := handleRes.Delete(userId)
	if err != nil {
		fmt.Println(err)
	}
	displayConsole.Line()
	fmt.Printf("The user with ID= %v was successfully deleted\n", userId)
	displayConsole.Line()
}

func getUsers(handleRes api.IApi) {
	result, err := handleRes.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	jsonStr, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		displayConsole.Line()
		fmt.Println(string(jsonStr))
		displayConsole.Line()
	}
}

func getStat(handleRes api.IApi) {
	result := handleRes.GetStatistics()
	jsonStr, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
