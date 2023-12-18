package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// note represents data about a note crated.
type note struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Description string  `json:"descriiption"`
    
}

// notes slice to seed record note data.
var notes = []note{
    {ID: "1", Title: "Simple Api", Description: "An api I wish to understand the principle"},
    {ID: "2", Title: "Server Backend", Description: "Nodejs backend"},
    {ID: "3", Title: "Cooking", Description: "Burger plus Fried Chicken"},
}

func main() {
    router := gin.Default()
    router.GET("/notes", getNotes)

    router.Run("localhost:8080")
}

// getNotes responds with the list of all notes as JSON.
func getNotes(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, notes)
}