package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	
)


// @BasePath /api/v1
// NoteExample godoc
// @Summary Get a list of tags
// @Schemes http https
// @Description get Notes
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /tags [get]
func getTags(c *gin.Context) {
	fmt.Println("getting Tags")
}


// @BasePath /api/v1
// NoteExample godoc
// @Summary Create a tag
// @Schemes http https
// @Description create a tag with provided content
// @Tags tags
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Router /tags [post]
func addTag(c *gin.Context) {
	fmt.Println("adding Tags")
}

// @BasePath /api/v1
// NoteExample godoc
// @Summary Get a specific tag
// @Schemes http https
// @Description get a tag based  on ID
// @Tags tags
// @Param tag body string true "Tag content"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /tags/:tagId [get]
func getTagById(c *gin.Context) {
	fmt.Println("Tag By Id")
}

// @BasePath /api/v1
// Note Example godoc
// @Summary Update a specific tag
// @Schemes http https
// @Description update a tag based  on ID
// @Tags tags
// @Param tag body string true "Tag content"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /tags/:tagId [put]
func updateTag(c *gin.Context) {
	fmt.Println("Updating a tag")
}


// @BasePath /api/v1
// NoteExample godoc
// @Summary Delete a specific tag
// @Schemes http https
// @Description delete a tag based on ID
// @Tags tags
// @Param tag body string true "Tag content"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /tags/:tagId [delete]
func deleteTag(c *gin.Context) {
	fmt.Println("Deleting a tag")
}



func addTagsRoutes(rg *gin.RouterGroup) {

	tags := rg.Group("/tags")

	tags.GET("/", getTags)
	tags.POST("/", addTag)
	tags.GET("/:tagId", getTagById)
	tags.PUT("/:tagId", updateTag)
	tags.DELETE("/:tagId", deleteTag)
}