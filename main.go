package main

import (
    "exercise/controllers/exercise"
    "fmt"
    "github.com/labstack/echo"
    "github.com/labstack/gommon/log"
    "gopkg.in/mgo.v2"
    "os"
    "time"
)

const (
    httpPort = 17000
)

func main() {
    _ = os.Setenv("TZ", "America/Montreal")

    app := echo.New()

    mongo := connectToMongoDB()
    _ = mongo

    exerciseController := exercise.NewExerciseController()
    apiAuth := app.Group("exercise")
    apiAuth.GET("/", exerciseController.GetExercises)

    log.Fatal(app.Start(fmt.Sprint(":", httpPort)))
}

func connectToMongoDB() *mgo.Session {
    var mongoDialInfo *mgo.DialInfo
    switch os.Getenv("APP_ENV") {
    default:
        mongoDialInfo = &mgo.DialInfo{
            Addrs:    []string{"exercise-db-mongodb.dev.svc.cluster.local:27017"},
            Timeout:  5 * time.Second,
            Database: "simulator",
        }
    }

    mongoSession, err := mgo.DialWithInfo(mongoDialInfo)
    if err != nil {
        log.Fatal("Could not open MongDB connection: ", err)
    }

    return mongoSession
}
