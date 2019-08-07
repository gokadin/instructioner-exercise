package exercise

import (
	"errors"
	"exercise/app/exercise/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database = "instructioner"
	collection = "exercise"
)

type Repository struct {
    mongo *mgo.Session
}

func NewRepository(mongo *mgo.Session) *Repository {
	return &Repository{
        mongo: mongo,
	}
}

func (r Repository) GetAll() ([]*model.Exercise, error) {
	var result = make([]*model.Exercise, 0)
	if err := r.mongo.DB(database).C(collection).Find(nil).All(&result); err != nil {
		return nil, errors.New("unable to retrieve camera list")
	}

	log.Debug("Fetched ", len(result), " exercises.")
	
	return result, nil
}

func (r Repository) Create(exercise *model.Exercise) error {
	exercise.Id = bson.NewObjectId()
	err := r.mongo.DB(database).C(collection).Insert(exercise)
	if err != nil {
		return errors.New(fmt.Sprint("Unable to insert exercise", err))
	}

	log.Debug("Added or updated exercise")

	return nil
}
