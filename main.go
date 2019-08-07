package main

import (
    "exercise/app/exercise"
    "exercise/config"
    "fmt"
    "github.com/jetbasrawi/go.geteventstore"
    "github.com/labstack/echo"
    "github.com/labstack/gommon/log"
    "github.com/tkanos/gonfig"
    "gopkg.in/mgo.v2"
    "os"
    "time"
)

const (
    httpPort = 17000
)

func main() {
	initialize()
	
	configuration := getConfiguration()

    mongo := connectToMongoDB(configuration)
    
    eventstore := connectToEventstore(configuration)
    _ = eventstore

    app := echo.New()
    setUpRoutes(app, mongo)
    log.Fatal(app.Start(fmt.Sprint(":", httpPort)))
}

func initialize() {
    _ = os.Setenv("TZ", "America/Montreal")
}

func getConfiguration() *config.Configuration {
    configuration := &config.Configuration{}
    env := os.Getenv("ENV")
    if len(env) == 0 {
        env = "development"
    }
    
    err := gonfig.GetConf(fmt.Sprintf("config/config.%s.json", env), configuration)
    if err != nil {
        log.Fatal("Could not get configuration.")
        os.Exit(500)
    }
    
    return configuration
}

func connectToMongoDB(config *config.Configuration) *mgo.Session {
    var mongoDialInfo *mgo.DialInfo
    switch os.Getenv("APP_ENV") {
    default:
        mongoDialInfo = &mgo.DialInfo{
            Addrs:    []string{fmt.Sprintf("%s:%d", config.MongodbHostname, config.MongodbPort)},
            Timeout:  5 * time.Second,
            Database: "simulator",
        }
    }

    mongoSession, err := mgo.DialWithInfo(mongoDialInfo)
    if err != nil {
        log.Fatal("Could not open MongDB connection: ", err)
        os.Exit(500)
    }

    return mongoSession
}

func connectToEventstore(config *config.Configuration) *goes.Client {
    client, err := goes.NewClient(nil, fmt.Sprintf("http://%s:%d", config.EventstoreHostname, config.EventstorePort))
    if err != nil {
        log.Fatal(err)
        os.Exit(500)
    }
    
    return client
}

func setUpRoutes(app *echo.Echo, mongo *mgo.Session) {
    exerciseController := exercise.NewController(mongo)
    apiAuth := app.Group("exercise")
    apiAuth.GET("", exerciseController.GetExercises)
    apiAuth.POST("/create", exerciseController.Create)
}