package main

import (
	"fmt"
	"postgre-project/common/env"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/repository"
)

func main() {

	// Database
	env.Load()
	database.ConnectDB()
	database.Instance.AutoMigrate(&model.Tables{})

	// SignUp
	repository.AddToDatabase(dto.DtoSignUp{
		Password:     "aaa",
		Token:        "",
		RefreshToken: "",
		FirstName:    "",
		LastName:     "",
		Email:        "aaa",
		UserType:     "",
	})

	// LogIn
	find, errFind := repository.FindInDatabase(dto.DtoLogIn{
		Email:    "aaa",
		Password: "aaa",
	})

	// GetResult
	resById, errId := repository.GetInfoByIdFromDatabase(5)

	// GetResults
	res, errRes := repository.GetInfosFromDatabase()

	// RESULTS
	pl := fmt.Printf

	pl("\nFinding: \n%v\n", find)
	pl("\nError: \n%v\n", errFind)
	pl("\nResult By ID: \n%v\n", resById)
	pl("\nError Result ID: \n%v\n", errId)
	// pl("\nResults: \n%v\n", res)
	for _, v := range res {
		pl("\nResults: \n%v   ---   %v\n", v.Email, v.Password)
	}
	pl("\nError Results: \n%v\n", errRes)

	database.CloseDB()
}