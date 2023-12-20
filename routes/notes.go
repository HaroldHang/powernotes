package routes

import (
	"fmt"
	"log"
	"net/http"
	"power_note/power_note_api/databases"
	"power_note/power_note_api/utils"
	"github.com/gin-gonic/gin"
	"encoding/json"
	
)
type noteReq struct {
	Title string `json:"title"`
	Content string `json:"content"`
	TagId int64 `json:"tagId"`
	UserId int64 `json:"userId"`

}

// @BasePath /api/v1

// NoteExample godoc
// @Summary note example
// @Schemes
// @Description do notes
// @Tags notes
// @Accept json
// @Produce json
// @Success 200 {json} Notes
// @Router /notes [get]
func getNotes(c *gin.Context) {
	notes, err := databases.GetNotes()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, notes)
}


func addNote(c *gin.Context) {
	var newNote noteReq
	if err := c.BindJSON(&newNote); err != nil {
		c.String(http.StatusOK, `the body should be reqA`)

		
	}
	note := databases.Note {
		UniqueId: utils.GenerateUUID(),
		Title: newNote.Title,
		Content: newNote.Content,
		Metadata: "",
	}
	noteId, err := databases.AddNote(note)
	if err != nil {
		fmt.Println("Not added")
	}
	note.Id = noteId
	noteData, err := json.Marshal(note)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(http.StatusOK, noteData)
}

func addNotesRoutes(rg *gin.RouterGroup) {

	notes := rg.Group("/notes")

	notes.GET("/", getNotes)
	notes.POST("/", addNote)

}
