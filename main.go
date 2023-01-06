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
	{
		Id:        "2",
		Title:     "Iron Man",
		BoxOffice: "300000",
		Director: &Director{
			Fullname: "John Doe",
		},
	},
}

func main() {
	router := gin.Default()

	router.GET("/movies", getAllMovies)
	router.POST("/movies", createMovie)
	router.GET("/movies/:id", getMovieById)
	router.PUT("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovieById)

	router.Run(":3000")
}

func deleteMovieById(context *gin.Context) {
	id := context.Param("id")
	for idx, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:idx], movies[idx+1:]...)
		}
	}
}

func updateMovie(context *gin.Context) {
	var newUpdatedMovie Movie

	if err := context.BindJSON(&newUpdatedMovie); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

func getMovieById(context *gin.Context) {
	id := context.Param("id")

	for _, movie := range movies {
		if movie.Id == id {
			context.IndentedJSON(http.StatusOK, movie)
			return
		}
	}
}

func createMovie(context *gin.Context) {
	var newMovie Movie

	if err := context.BindJSON(&newMovie); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	movies = append(movies, newMovie)
	context.IndentedJSON(http.StatusCreated, newMovie)
}

func getAllMovies(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, movies)
}
