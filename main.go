package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type UserPageSite struct {
	Username string `json:"username"`
	Site     string `json:"site"`
}

func userExist(dirs []fs.FileInfo, name string) bool {
	for _, f := range dirs {
		if f.Name() == name {
			return true
		}
	}
	return false
}

func getUserPageHandler(c *gin.Context) {
	name := c.Param("name")
	files, err := ioutil.ReadDir("./Database")
	if err != nil {
		log.Fatal(err)
	}

	if !userExist(files, name) {
		c.Status(http.StatusNotFound)
		return
	}
	var data []UserPageSite
	jsonFile, err := ioutil.ReadFile("./Database/" + name + "/data.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, data)

}

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/user/:name", getUserPageHandler)

	r.Run(":8080")

}
