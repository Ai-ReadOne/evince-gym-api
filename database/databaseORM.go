package database

import (
	"fmt"
	"time"

	postgres "github.com/go-pg/pg"
	"github.com/google/uuid"
)

// creating an object for the gym customers
type Person struct {
	LastName  string    `sql:",lastname,type:varchar(21),notnull"`
	FirstName string    `sql:",firstname,type:varchar(21),notnull"`
	Gender    string    `sql:",gender,notnull"`
	Age       int       `sql:",age,notnull"`
	Phone     uint      `sql:",phone,type:int(10),Unique,notnull"`
	Email     string    `sql:",email,type:varchar(35),Unique,notnull"`
	UpdatedAt time.Time `sql:",updated_at,notnull,default time.Now()"`
}

type GymMember struct {
	Person
	MemberID           uuid.UUID `sql:",member_id,Unique,primary_key,type:uuid,default uuid.New(),fk:instructor_id"`
	InstructorID       uuid.UUID `sql:",instructor_id,primary_key"`
	JoinDate           time.Time `sql:",join_date,notnull,default time.Now()"`
	LastSeen           time.Time `sql:",lastseen,"`
	MembershipValidity string    `sql:",membership_validity,"`
}

// creating an object for the gym instructors
type GymInstructor struct {
	Person
	InstructorID uuid.UUID `sql:",instructor_id, unique, pk, type:uuid, default uuid.New()"`
	EmpDate      time.Time `sql:",emp_date, default time.Now(), notnull"`
}

func (member *GymMember) SaveNewMember(database *postgres.DB) error {
	member.UpdatedAt = time.Now()
	member.MemberID = uuid.New()
	member.JoinDate = time.Now()
	member.MembershipValidity = time.Now().Format("2022-01-30")
	insertError := database.Insert(member)
	if insertError != nil {
		fmt.Printf("Error while inserting new record to database %v\n", insertError)
	} else {

		fmt.Println("new In structor inserted sucessfully")

	}

	return insertError
}

func (instructor *GymInstructor) SaveNewInstructor(database *postgres.DB) error {
	instructor.UpdatedAt = time.Now()
	instructor.EmpDate = time.Now()
	instructor.InstructorID = uuid.New()

	insertError := database.Insert(instructor)
	if insertError != nil {
		fmt.Printf("Error while inserting new record to database %v\n", insertError)
		return insertError
	} else {

		fmt.Println("New Instructor inserted sucessfully")

	}
	return nil
}

func FetchAllInstructor(database *postgres.DB) ([]GymInstructor, error) {
	var instructors []GymInstructor
	getError := database.Model(&instructors).Select()
	if getError != nil {
		fmt.Println("Error while Fecthing data from databse")
	}
	fmt.Print(instructors)
	return instructors, getError
}

// func (member *GymMember) GetAllMembers(database *postgres.DB) error {
// 	getError := database.Select(member)
// 	if getError == nil {
// 		fmt.Printf("Error while Fecthing data from databse \n")
// 	}

// 	return getError
// }
