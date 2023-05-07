package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func carById(c *gin.Context) {
	id := c.Param("id")
	car, err := getCarById(id)

	if err != nil {
		// handling response in case of error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, car)

}

func getCarById(id string) (*car, error) {

	for i, c := range cars {

		if c.ID == id {
			return &cars[i], nil
		}
	}

	return nil, errors.New("Car not found")
}

func buyCar(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id - query parameter"})
		return
	}

	car, err := getCarById(id)

	if err != nil {
		// handling response in case of error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
		return
	}

	if car.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Car not available"})
		return
	}

	car.Quantity -= 1

	c.IndentedJSON(http.StatusOK, car)

}

func sellCar(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id - query parameter"})
		return
	}

	car, err := getCarById(id)

	if err != nil {
		// handling response in case of error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
		return
	}

	car.Quantity += 1

	c.IndentedJSON(http.StatusOK, car)

}

func main() {
	fmt.Println("Main Function of Car API")

	router := gin.Default() // creating router object
	router.GET("/cars", getCars)
	router.GET("cars/:id", carById) // added path parameter id
	router.POST("/cars", addNewCar)
	router.PUT("/buycar", buyCar)   // Using Query parameter in request
	router.PUT("/sellcar", sellCar) //Using Query parameter in request
	router.Run("localhost:8080")
}
