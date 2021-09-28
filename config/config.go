package config

import (
	"github.com/joho/godotenv"
	"github.com/klovercloud-ci/enums"
	"log"
	"os"
)

var ServerPort string
var DbServer string
var DbPort string
var DbUsername string
var DbPassword string
var DatabaseConnectionString string
var DatabaseName string
var Database string
var KlovercloudCiCoreUrl string
var KlovercloudCiCoreToken string
var EventStoreUrl string
var EventStoreToken string

func InitEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
	ServerPort = os.Getenv("SERVER_PORT")
	DbServer = os.Getenv("MONGO_SERVER")
	DbPort = os.Getenv("MONGO_PORT")
	DbUsername = os.Getenv("MONGO_USERNAME")
	DbPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	EventStoreUrl = os.Getenv("EVENT_STORE_URL")
	EventStoreToken = os.Getenv("EVENT_STORE_TOKEN")
	Database = os.Getenv("DATABASE")
	if Database == enums.Mongo {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}
	KlovercloudCiCoreUrl = os.Getenv("KLOVERCLOUD_CI_CORE_URL")
	KlovercloudCiCoreToken = os.Getenv("KLOVERCLOUD_CI_CORE_TOKEN")
}
