package config

import (
	"github.com/joho/godotenv"
	"github.com/klovercloud-ci/enums"
	"log"
	"os"
	"strconv"
	"strings"
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
var EventStoreUrl string
var Publickey string
var EnableAuthentication bool
var Token string
var DefaultPerDayTotalBuild,DefaultNumberOfConcurrentBuild int64
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
	Database = os.Getenv("DATABASE")
	if Database == enums.Mongo {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}
	KlovercloudCiCoreUrl = os.Getenv("KLOVERCLOUD_CI_CORE_URL")
	if os.Getenv("ENABLE_AUTHENTICATION") == "" {
		EnableAuthentication = false
	} else {
		if strings.ToLower(os.Getenv("ENABLE_AUTHENTICATION")) == "true" {
			EnableAuthentication = true
		} else {
			EnableAuthentication = false
		}
	}
	Token = os.Getenv("TOKEN")

	DefaultPerDayTotalBuild, err = strconv.ParseInt(os.Getenv("DEFAULT_PER_DAY_TOTAL_BUILD"), 10, 64)
	if err!=nil{
		DefaultPerDayTotalBuild=10
	}
	DefaultNumberOfConcurrentBuild, err = strconv.ParseInt(os.Getenv("DEFAULT_NUMBER_OF_CONCURRENT_BUILD"), 10, 64)
	if err!=nil{
		DefaultNumberOfConcurrentBuild=10
	}
}
