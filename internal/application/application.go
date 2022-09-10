package application

import (
	"net/http"
	"os"

	"github.com/bsgilber/heart-beat/internal/controllers"
	"github.com/bsgilber/heart-beat/internal/databases"
	"github.com/gin-gonic/gin"
)

var environment string

type App struct {
	srv *gin.Engine
}

func New() *App {
	InitEnv()
	application := App{}

	db := databases.ConnectClient()
	endpointDb := databases.NewEndpointDb(db)
	h := controllers.NewBaseHandler(endpointDb)

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/pingall", h.AllPingPrep)
	// This handler will match /ping/john but will not match /ping/ or /ping
	router.GET("/ping/:name", h.SinglePingPrep)
	router.GET("/health", health)
	router.GET("/list", h.ListRegistered)
	// Example for binding JSON ({"name": "unique_name", "url": "https://some.url.com"})
	router.POST("/register", h.Register)

	application.srv = router

	return &application
}

func (a *App) Start() {
	a.srv.Run(":8080")
}

func InitEnv() {
	val, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		val = "local"
	}

	environment = val
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
