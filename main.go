package main

import (
	"fmt"
	//"net/http"
	//"github.com/gin-gonic/gin"
	"power_note/power_note_api/databases"
	"power_note/power_note_api/routes"
)



// notes slice to seed record note data.


func main() {
	fmt.Println(databases.GetDB())
    //router := gin.Default()
    //router.GET("/notes", getNotes)

    //router.Run("localhost:8080")
    ATest()

    routes.Run()
}