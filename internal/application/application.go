package application

import (
	json "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var cache map[string]string
var environment string

type App struct {
	srv *gin.Engine
}

var Cache interface {
	Set(key string, data interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
}

// Binding from JSON
type AddEndpoint struct {
	Name string `form:"name" json:"name" binding:"required"`
	Url  string `form:"url" json:"url" binding:"required"`
}

func New() *App {
	InitLocalCache()
	InitEnv()

	application := App{}

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/pingall", allPingPrep)
	// This handler will match /ping/john but will not match /ping/ or /ping
	router.GET("/ping/:name", singlePingPrep)
	router.GET("/health", health)
	router.GET("/list", listRegistered)
	// Example for binding JSON ({"name": "unique_name", "url": "https://some.url.com"})
	router.POST("/register", register)

	application.srv = router

	return &application
}

func (a *App) Start() {
	a.srv.Run(":8080")
}

func InitLocalCache() {
	cache = map[string]string{}
}

func InitEnv() {
	val, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		val = "local"
	}

	environment = val
}

func listRegistered(c *gin.Context) {
	b, err := json.Marshal(cache)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, string(b))
}

func singlePingPrep(c *gin.Context) {
	name := c.Param("name")

	ping(c, name)
}

func allPingPrep(c *gin.Context) {
	for key := range cache {
		ping(c, key)
	}
}

func ping(c *gin.Context, name string) {
	if _, ok := cache[name]; !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("the key [%s] does not exist.\n", name))
		return
	}

	req, err := http.NewRequest(http.MethodGet, cache[name], nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	defer res.Body.Close()

	c.String(http.StatusOK, "[%s] had status code [%d]\n", name, res.StatusCode)
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}

func register(c *gin.Context) {
	var json AddEndpoint

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := cache[json.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "name must be unique"})
		return
	}

	u, err := url.Parse(json.Url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		c.String(http.StatusBadRequest, "invalid url, url must contain scheme, host, and/or path")
		return
	}

	cache[json.Name] = json.Url
	c.String(http.StatusOK, fmt.Sprintf("[%s] has been added to the cache.", json.Name))
}
