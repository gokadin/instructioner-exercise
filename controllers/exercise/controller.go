package exercise

import (
	"exercise/models"
	"github.com/labstack/echo"
	"net/http"
)

type Exercise struct {
}

func NewExerciseController() *Exercise {
	return &Exercise{}
}

func (e Exercise) GetExercises(c echo.Context) error {
    var exercises = []models.Exercise{
		{Id: 1, Name: "one"},
		{Id: 2, Name: "two"},
    }
    
    c.JSON(http.StatusOK, exercises)
    
    return nil
}
