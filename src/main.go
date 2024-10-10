package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type app_info struct {
	Name string `json:"name"`
	Env  string `json:"env"`
	UUID string `json:"uuid"`
}

func getAppInfo(c *gin.Context) {
	uuid := "123123123"
	// if UUID is defined as an environment variable, replace the value
	if nuuid := os.Getenv("UUID"); nuuid != "" {
		uuid = nuuid
	}
	var info = app_info{
		Name: "echo-app",
		Env:  "dev",
		UUID: uuid,
	}
	c.IndentedJSON(http.StatusOK, &info)
}

func main() {
	router := gin.Default()
	router.GET("/", getAppInfo)
	router.Run(":8889")
}
