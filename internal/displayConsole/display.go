package displayConsole

import (
	"fmt"
	"github.com/Shemistan/Lesson_5/internal/models"
	"log"
	"time"
)

func RegistrationDataFromConsole() models.User {
	var initialData models.User
	fmt.Println("Please enter login")
	_, err := fmt.Scan(&initialData.Login)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Please enter password")
	_, err = fmt.Scan(&initialData.PasswordHash)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Please enter name")
	_, err = fmt.Scan(&initialData.Name)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Please enter surname")
	_, err = fmt.Scan(&initialData.Surname)
	if err != nil {
		log.Print(err)
	}
	return initialData
}

func GetUserIdFromConsole() int {
	var userId int
	fmt.Println("Please enter user ID")
	_, err := fmt.Scan(&userId)
	if err != nil {
		log.Print(err)
	}
	return userId
}

func GetDataFromConsoleToUpdate() (int, *models.UserDate) {
	var userId int
	var fromInput string
	var userData models.UserDate
	fmt.Println("Please, enter user ID")
	_, err := fmt.Scan(&userId)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Please enter name")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		log.Print(err)
	}
	userData.Name = fromInput

	fmt.Println("Please enter surname")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		log.Print(err)
	}
	userData.Surname = fromInput

	fmt.Println("Please enter status")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		log.Print(err)
	}
	userData.Status = fromInput

	fmt.Println("Please enter role")
	_, err = fmt.Scan(&fromInput)
	if err != nil {
		log.Print(err)
	}
	userData.Role = fromInput
	userData.UpdateDate = time.Now().Format("01-02-2006 15:04:05")

	return userId, &userData
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

//log.Printf("user %v is added: %+v", s.ids, user)
//jsonStr, err := json.MarshalIndent(s.db, "", "\t")
//if err != nil {
//fmt.Printf("Error: %s", err.Error())
//} else {
//fmt.Println(string(jsonStr))
//}
//fmt.Printf("%v", s.db)

//jsonStr, err := json.MarshalIndent(data, "", "\t")
//if err != nil {
//return "", err
//}
//fmt.Println("%T", jsonStr)
//return string(jsonStr), nil

//jsonStr, err := json.MarshalIndent(gotUser, "", "\t")
//if err != nil {
//fmt.Printf("Error: %s", err.Error())
//} else {
//fmt.Println(string(jsonStr))
//}
