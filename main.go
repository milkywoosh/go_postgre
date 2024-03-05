package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// get route
func fetchAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// post route
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func fetchAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// IN REAL IMPLEMENTATION, MUST USE DB TO DO A QUERY find by ID
	for _, val := range albums {
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album is not found"})
}

func main() {
	fmt.Println("test")

	var router *gin.Engine
	router = gin.Default()
	router.GET("/albums", fetchAlbums)
	router.POST("/albums/create-new", postAlbums)
	router.GET("/albums/:id", fetchAlbumsByID)

	log.Fatal(router.Run("localhost:8080"))
}
