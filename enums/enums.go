package enums

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

// REPOSITORY_TYPE repository types[may be any git]
type REPOSITORY_TYPE string

const (
	// GITHUB github as repository
	GITHUB = REPOSITORY_TYPE("GITHUB")
	// BIT_BUCKET bitbucket as repository
	BIT_BUCKET = REPOSITORY_TYPE("BIT_BUCKET")
)

// COMPANY_UPDATE_OPTION company update options
type COMPANY_UPDATE_OPTION string

const (
	// APPEND_APPLICATION company update option to append application
	APPEND_APPLICATION = COMPANY_UPDATE_OPTION("APPEND_APPLICATION")
	// APPEND_REPOSITORY company update option to append repository
	APPEND_REPOSITORY = COMPANY_UPDATE_OPTION("APPEND_REPOSITORY")
	// SOFT_DELETE_APPLICATION company update option to soft delete application
	SOFT_DELETE_APPLICATION = COMPANY_UPDATE_OPTION("SOFT_DELETE_APPLICATION")
	// DELETE_APPLICATION company update option to delete application
	DELETE_APPLICATION = COMPANY_UPDATE_OPTION("DELETE_APPLICATION")
	// SOFT_DELETE_REPOSITORY company update option to soft delete repository
	SOFT_DELETE_REPOSITORY = COMPANY_UPDATE_OPTION("SOFT_DELETE_REPOSITORY")
	// DELETE_REPOSITORY company update option to delete repository
	DELETE_REPOSITORY = COMPANY_UPDATE_OPTION("DELETE_REPOSITORY")
)

// COMPANY_STATUS company status options
type COMPANY_STATUS string

const (
	// ACTIVE company status for active company
	ACTIVE = COMPANY_STATUS("ACTIVE")
	// INACTIVE company status for inactive company
	INACTIVE = COMPANY_STATUS("INACTIVE")
)

// STEP_TYPE steps type
type STEP_TYPE string

const (
	// BUILD build step
	BUILD = STEP_TYPE("BUILD")
	// DEPLOY deploy step
	DEPLOY = STEP_TYPE("DEPLOY")
)

// GITHUB_URL gitbhub url for different operations
type GITHUB_URL string

const (
	// GITHUB_RAW_CONTENT_BASE_URL gitbhub url for raw content
	GITHUB_RAW_CONTENT_BASE_URL = "https://raw.githubusercontent.com/"
	// GITHUB_BASE_URL gitbhub base url
	GITHUB_BASE_URL = "https://github.com/"
	// GITHUB_API_BASE_URL gitbhub base url for api access
	GITHUB_API_BASE_URL = "https://api.github.com/"
)

const (
	// PIPELINE_FILE_NAME pipeline containing file name
	PIPELINE_FILE_NAME = "pipeline"
)

// TRIGGER pipeline trigger options
type TRIGGER string

const (
	// AUTO pipeline trigger options is auto
	AUTO = TRIGGER("AUTO")
	// MANUAL pipeline trigger options is MANUAL
	MANUAL = TRIGGER("MANUAL")
)

// PARAMS pipeline parameters
type PARAMS string

const (
	// REVISION resource revision key for  pipeline step param
	REVISION = PARAMS("revision")
)

// PIPELINE_FILE_BASE_DIRECTORY pipeline file base directory
const PIPELINE_FILE_BASE_DIRECTORY = "klovercloud/pipeline"

// PIPELINE_DESCRIPTORS_BASE_DIRECTORY pipeline descriptors base directory
const PIPELINE_DESCRIPTORS_BASE_DIRECTORY = "klovercloud/pipeline/configs"

// GIT_EVENT git web hook event options
type GIT_EVENT string

const (
	// PUSH git web hook push event option
	PUSH = GIT_EVENT("push")
	// RELEASE git web hook release event option
	RELEASE = GIT_EVENT("release")
	// DELETE git web hook delete event option
	DELETE = GIT_EVENT("delete")
)
