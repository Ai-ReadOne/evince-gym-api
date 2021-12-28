package database

import (
	"fmt"
	"time"

	postgres "github.com/go-pg/pg"
	"github.com/google/uuid"
)

// creating an object for the gym customers
type Person struct {
	LastName  string    `sql:",lastname,type:varchar(21),notnull" validate:"required,max=21,alpha"`
	FirstName string    `sql:",firstname,type:varchar(21),notnull" validate:"required,alpha,max=21"`
	Gender    string    `sql:",gender,notnull" validate:"required,alpha,max=1"`
	Age       int       `sql:",age,notnull" validate:"gte=15,required"`
	Phone     uint      `sql:",phone,type:int(10),unique,notnull" validate:"required,min=10,numeric"`
	Email     string    `sql:",email,type:varchar(35),unique,notnull" validate:"required,max=35,email"`
	UpdatedAt time.Time `sql:",updated_at,notnull" validate:"omitempty"`
}

// creating an object for the gym Members
type GymMember struct {
	Person
	MemberID           uuid.UUID `sql:",member_id,unique,primary_key,type:uuid,fk:instructor_id" validate:"required"`
	InstructorID       uuid.UUID `sql:",instructor_id,primary_key" validate:"required"`
	JoinDate           time.Time `sql:",join_date,notnull" validate:"required"`
	LastSeen           time.Time `sql:",lastseen," validate:"required"`
	MembershipValidity string    `sql:",membership_validity," validate:"required"`
}

// creating an object for the gym instructors
type GymInstructor struct {
	Person
	InstructorID uuid.UUID `sql:",instructor_id, unique, pk, type:uuid" validate:"omitempty"`
	EmpDate      time.Time `sql:",emp_date, notnull" validate:"omitempty"`
}

// Create new member
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

// creating new instructor
func (instructor *GymInstructor) SaveNewInstructor(database *postgres.DB) error {
	instructor.UpdatedAt = time.Now()
	instructor.EmpDate = time.Now()
	instructor.InstructorID = uuid.New()

	insertError := database.Insert(instructor)
	if insertError != nil {
		fmt.Printf("Error while inserting new record to database %v\n", insertError)
		return insertError
	}
	fmt.Println("New Instructor inserted sucessfully")

	return insertError
}

// fecthing all instructors from database
func FetchAllInstructor(database *postgres.DB) ([]GymInstructor, error) {
	var instructors []GymInstructor
	getError := database.Model(&instructors).Select()
	if getError != nil {
		fmt.Println("Error while Fecthing data from databse")
	}
	fmt.Print(instructors)
	return instructors, getError
}

// fetching all gym members from database
func (member *GymMember) GetAllMembers(database *postgres.DB) ([]GymMember, error) {
	var members []GymMember
	getError := database.Model(&members).Select()
	if getError != nil {
		fmt.Println("Error while Fecthing data from databse")
	}
	fmt.Print(members)
	return members, getError
}
