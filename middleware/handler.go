package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandleCreateInstructor() gin.HandlerFunc {
	// Instructor := database.GymInstructor{}
	// validate := validator.New()
	// Instructor.EmpDate = time.Now()
	// Instructor.UpdatedAt = time.Now()
	// Instructor.InstructorID = uuid.New()

	return func(wr *gin.Context) {
		// if error := wr.ShouldBindBodyWith(&Instructor, binding.JSON); error != nil {
		// 	wr.JSON(http.StatusBadRequest, gin.H{
		// 		"Error21": error.Error(),
		// 	})
		// }
		// BodyBytes, _ := ioutil.ReadAll(wr.Request.Body)
		// wr.Request.Body = ioutil.NopCloser(bytes.NewBuffer(BodyBytes))
		// fmt.Println(string(BodyBytes))

		// if error := validate.Struct(Instructor); error != nil {
		// 	wr.JSON(http.StatusBadRequest, gin.H{
		// 		"error1": error,
		// 	})
		// 	return
		// }

		// var BodyBytes []byte

		wr.Next()
	}
}
