package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Movie struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	BoxOffice string    `json:"boxOffice"`
	Director  *Director `json:"director"`
}

type Director struct {
	Fullname string `json:"fullname"`
}

var movies = []Movie{
	{
		Id:        "1",
		Title:     "Spider Man",
		BoxOffice: "3000000",
		Director: &Director{
			Fullname: "Aung Myat Moe",
		},
	},
}

func main() {
	router := gin.Default()

	router.GET("/movies", getAllMovies)

	router.Run(":3000")
}

func getAllMovies(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, movies)
}
