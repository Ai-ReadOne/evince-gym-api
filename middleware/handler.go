package middleWare

import (
	"evince-gym-api/database"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleCreateInstructor() gin.HandlerFunc {

	return func(wr *gin.Context) {
		instructor := database.GymInstructor{}
		instructor.EmpDate = time.Now()
		instructor.UpdatedAt = time.Now()
		instructor.InstructorID = uuid.New()
		fmt.Println(*wr)
	}
}
