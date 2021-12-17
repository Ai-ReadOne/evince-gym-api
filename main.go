package main

import (
	"log"

	// database "./database"

	"github.com/gin-gonic/gin"
	postgres "github.com/go-pg/pg"

	// "golang.org/x/text/date"
	// "gorm.io.gorm"

	"evince-gym-api/database"
	"evince-gym-api/middleWare"
	"evince-gym-api/routes"
)

// creating a record object that stores all the instructors and gym customers
func databaseServer(dbconnect *postgres.DB) gin.HandlerFunc {
	return func(wr *gin.Context) {
		wr.Set("database", dbconnect)
		wr.Next()
	}
}

func main() {
	dbConnect := database.Connect()

	Server := gin.Default()
	// fmt.Println(uuid.NewString())
	Server.Use(databaseServer(dbConnect))

	Server.GET("/", routes.Welcome)
	instructorParser := Server.Group("instructor")
	{

		instructorParser.POST("create", middleWare.HandleCreateInstructor(), routes.CreateNewInstructor)
		instructorParser.GET("get", routes.GetAllInstructor)
		instructorParser.GET("get/:InstructorID", routes.GetInstructorByID)
		instructorParser.PUT("update", routes.UpdateInstructor)

	}

	Server.Run(":8000")

	log.Println("This api is running on 127.0.0.1:8000")
}
