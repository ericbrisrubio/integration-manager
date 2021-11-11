package config

import (
	"github.com/joho/godotenv"
	"github.com/klovercloud-ci/enums"
	"log"
	"os"
	"strconv"
	"strings"
)

// ServerPort refers to server port.
var ServerPort string

// DbServer refers to database server ip.
var DbServer string

// DbPort refers to database server port.
var DbPort string

// DbUsername refers to database name.
var DbUsername string

// DbPassword refers to database password.
var DbPassword string

// DatabaseConnectionString refers to database connection string.
var DatabaseConnectionString string

// DatabaseName refers to database name.
var DatabaseName string

// Database refers to database options.
var Database string

// KlovercloudCiCoreUrl refers to ci core url.
var KlovercloudCiCoreUrl string

// EventStoreUrl refers to event-store url.
var EventStoreUrl string

// PublicKey refers to public key of EventStoreToken.
var PublicKey string

// EnableAuthentication refers if service to service authentication is enabled.
var EnableAuthentication bool

// Token refers to jwt token for service to service communication.
var Token string

// DefaultNumberOfConcurrentProcess refers to default number of concurrent process
var DefaultNumberOfConcurrentProcess int64

// DefaultPerDayTotalProcess refers to default number of build per day
var DefaultPerDayTotalProcess int64

// GithubWebhookConsumingUrl refers to github web hook consuming url.
var GithubWebhookConsumingUrl string

// InitEnvironmentVariables initializes environment variables
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
	if Database == enums.MONGO {
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

	DefaultPerDayTotalProcess, err = strconv.ParseInt(os.Getenv("DEFAULT_PER_DAY_TOTAL_PROCESS"), 10, 64)
	if err != nil {
		DefaultPerDayTotalProcess = 10
	}
	DefaultNumberOfConcurrentProcess, err = strconv.ParseInt(os.Getenv("DEFAULT_NUMBER_OF_CONCURRENT_PROCESS"), 10, 64)
	if err != nil {
		DefaultNumberOfConcurrentProcess = 10
	}
	GithubWebhookConsumingUrl = os.Getenv("GITHUB_WEBHOOK_CONSUMING_URL")
}
