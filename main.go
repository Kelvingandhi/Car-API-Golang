package main

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Quantity int    `json:"quantity"`
}

// Create cars slice to initialize with few values
var cars = []car{
	{ID: "1", Name: "Mercedes C-Class", Color: "Black", Quantity: 2},
	{ID: "2", Name: "Honda Civics", Color: "Silver", Quantity: 3},
	{ID: "3", Name: "GMC Hummer EV", Color: "White", Quantity: 4},
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func addNewCar(c *gin.Context) {
	var newCar car // create a new car var

	if err := c.BindJSON(&newCar); err != nil && err.Error() == "bind: address already in use" { // bind request body to newCar var
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, cars)

}

func main() {
	fmt.Println("Main Function of Car API")

	router := gin.Default() // creating router object
	router.GET("/cars", getCars)
	router.POST("/cars", addNewCar)
	router.Run("localhost:8080")
}
