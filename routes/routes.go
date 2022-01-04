package routes

import (
	"evince-gym-api/database"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	postgres "github.com/go-pg/pg"
	"github.com/go-playground/validator/v10"
)

func Welcome(wr *gin.Context) {
	wr.JSON(http.StatusOK, "Welcome to the evince system Gym application")
}

func CreateNewInstructor(wr *gin.Context) {
	Instructor := database.GymInstructor{}
	if error := wr.ShouldBindBodyWith(&Instructor, binding.JSON); error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"Error1": error.Error(),
		})
		return
	}

	validate := validator.New()
	if error := validate.Struct(Instructor); error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"error12": error.Error(),
		})
		return
	}

	databaseConnect, _ := wr.Get("database")
	connect := databaseConnect.(*postgres.DB)
	if error := Instructor.SaveNewInstructor(connect); error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"Error123": error.Error(),
		})
		return
	}

	wr.JSON(http.StatusOK, gin.H{
		"message":       "new Instructor Created sucessfully",
		"instructor_id": Instructor.InstructorID,
		"LastName":      Instructor.LastName,
		"FirstName":     Instructor.FirstName,
		"Empdate":       Instructor.EmpDate,
		"UpdatedAt":     Instructor.UpdatedAt,
	})

}

func GetAllInstructor(wr *gin.Context) {
	// var Instructor database.GymInstructor

	databaseConnect, _ := wr.Get("database")
	connect := databaseConnect.(*postgres.DB)

	Instructors, error := database.FetchAllInstructor(connect)
	if error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"Error1": error.Error(),
		})
		return
	}

	wr.JSON(http.StatusOK, gin.H{
		"data": Instructors,
	})
}

func GetInstructorByID(wr *gin.Context) {

}

func UpdateInstructor(wr *gin.Context) {

}

func CreateNewMember(wr *gin.Context) {
	Member := database.GymMember{}
	if error := wr.ShouldBindBodyWith(&Member, binding.JSON); error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"Error1": error.Error(),
		})
		return
	}

	validate := validator.New()
	if error := validate.Struct(Member); error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"error12": error.Error(),
		})
		return
	}

	databaseConnect, _ := wr.Get("database")
	connect := databaseConnect.(*postgres.DB)
	if error := Member.SaveNewMember(connect); error != nil {
		wr.JSON(http.StatusBadRequest, gin.H{
			"Error123": error.Error(),
		})
		return
	}
	wr.JSON(http.StatusOK, gin.H{
		"message":       "new Instructor Created sucessfully",
		"instructor_id": Member.InstructorID,
		"member_id":     Member.MemberID,
		"LastName":      Member.LastName,
		"FirstName":     Member.FirstName,
	})

}
