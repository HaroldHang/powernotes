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
// @Summary Get a list of notes
// @Schemes http https
// @Description get Notes
// @Tags notes
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /notes [get]
func getNotes(c *gin.Context) {
	notes, err := databases.GetNotes()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, notes)
}


// @BasePath /api/v1
// NoteExample godoc
// @Summary Create a note
// @Schemes http https
// @Description create a note with provided content
// @Tags notes
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Router /notes [post]
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

// @BasePath /api/v1
// NoteExample godoc
// @Summary Get a specific note
// @Schemes http https
// @Description get a note based  on ID
// @Tags notes
// @Param note body string true "Note content"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /notes/:noteId [get]
func getNoteById(c *gin.Context) {
	fmt.Println("Note By Id")
}

// @BasePath /api/v1
// NoteExample godoc
// @Summary Update a specific note
// @Schemes http https
// @Description update a note based  on ID
// @Tags notes
// @Param note body string true "Note content"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /notes/:noteId [put]
func updateNote(c *gin.Context) {
	fmt.Println("Updating a note")
}


// @BasePath /api/v1
// NoteExample godoc
// @Summary Delete a specific note
// @Schemes http https
// @Description delete a note based on ID
// @Tags notes
// @Param note body string true "Note content"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /notes/:noteId [delete]
func deleteNote(c *gin.Context) {
	fmt.Println("Deleting a note")
}

func addNotesRoutes(rg *gin.RouterGroup) {

	notes := rg.Group("/notes")

	notes.GET("/", getNotes)
	notes.POST("/", addNote)
	notes.GET("/:noteId", getNoteById)
	notes.PUT("/:noteId", updateNote)
	notes.DELETE("/:noteId", deleteNote)
}
