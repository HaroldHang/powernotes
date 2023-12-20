package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	docs "power_note/power_note_api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"html/template"
	"net/http"
	"embed"
)

var router = gin.Default()
var f embed.FS

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy

// @title Power Notes API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1

func getRoutes() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("api/v1")
	addNotesRoutes(v1)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//v2 := router.Group("/v2")
	//addPingRoutes(v2)
}

// Run will start the server
func Run() {
	//templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	//router.SetHTMLTemplate(templ)
	router.LoadHTMLGlob("templates/*")
	// example: /public/assets/images/example.png
	router.StaticFS("/public", http.FS(f))
	router.GET("/", renderStatic)
	router.GET("/:uri", renderStatic)
	getRoutes()
	router.Run(":8080")
}

func renderStatic(c *gin.Context) {
	temp := c.Param("uri")
	if temp == "" {
		temp = "index"
	}
	c.HTML(http.StatusOK, fmt.Sprintf("%s.tmpl", temp), gin.H{
		"title": fmt.Sprintf("%s Page", temp),
	})
}