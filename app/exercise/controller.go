package exercise

import (
	"encoding/json"
	"exercise/app/exercise/model"
	"exercise/app/exercise/parser"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"net/http"
)

type Controller struct {
	repository *Repository
}

func NewController(mongo *mgo.Session) *Controller {
	return &Controller{
		repository: NewRepository(mongo),
	}
}

func (e Controller) GetExercises(c echo.Context) error {
    exercises, err := e.repository.GetAll()
    if err != nil {
        log.Error("Could not retrieve exercises", err)
        return err
	}
    
    for _, exercise := range exercises {
        err := parser.Parse(exercise)
        if err != nil {
        	log.Error("Could not parse exercise", err)
        	return err
		}
	}
    
    c.JSON(http.StatusOK, exercises)
    
    return nil
}

func (e Controller) Create(c echo.Context) error {
    exercise := &model.Exercise{}
    err := json.NewDecoder(c.Request().Body).Decode(&exercise)
    if err != nil {
		log.Error("Could not decode exercise", err)
    	return err
	}
    
    err = e.repository.Create(exercise)
    if err != nil {
    	log.Error("Could not create exercise", err)
        return err
	}

	c.JSON(http.StatusOK, nil)

	return nil
}