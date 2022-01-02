package database

import (
	"fmt"

	postgres "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

func Connect() *postgres.DB {
	options := &postgres.Options{
		User:     "postgres",
		Password: "aireadone",
		Addr:     "localhost:5432",
	}

	database := postgres.Connect(options)
	if database == nil {
		fmt.Println("unable to connect to database")
	} else {
		fmt.Println("database connected succesfully")
		fmt.Println(database)
	}
	CreateInstructorTable(database)
	CreateMemberTable(database)

	return database

}

func CreateMemberTable(database *postgres.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := database.CreateTable(&GymMember{}, options)
	if createError != nil {
		fmt.Printf("Error 1 occured while creating Table %v\n", createError.Error())
		return createError
	} else {
		fmt.Println(" Gym membersTable created Succesfully")
	}
	return createError
}

func CreateInstructorTable(database *postgres.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := database.CreateTable(&GymInstructor{}, options)
	if createError != nil {
		fmt.Printf("Error 2 occured while creating Table %v\n", createError)
		return createError
	} else {
		fmt.Println("Gym instructors Table created Succesfully")
	}

	return createError
}
