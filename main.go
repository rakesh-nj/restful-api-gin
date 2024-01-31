package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
type album struct {
    ID string `json:"id"`

    Title string `json:"id"`
    Artist string `json:"id"`
    Price float64 `json:"price"`
}

var albums = []album {
    {ID:"1",Title:"Title 1",Artist:"Artist 1",Price:10.99},
    {ID:"2",Title:"Title 2",Artist:"Artist 2",Price:20.99},
    {ID:"3",Title:"Title 3",Artist:"Artist 3",Price:30.99},
}

func getAlbums(c *gin.Context) {
   c.IndentedJSON(http.StatusOK,albums)
}



func main() {
    router := gin.Default()
    router.GET("/albums",getAlbums)
    router.GET("/albums/:id",getAlbumByID)
    router.POST("/albums",postAlbums)

    router.Run()
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    if err := c.BindJSON(&newAlbum); err != nil {
        return 
    }

    albums = append(albums,newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _,a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a) 
            return 
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}
