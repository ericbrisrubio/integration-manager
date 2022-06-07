package v1

// Agent contains payload data
type Agent struct {
	Agent      string `bson:"agent" json:"agent"`
	ApiVersion string `bson:"api_version" json:"api_version"`
	Terminal   string `bson:"terminal" json:"terminal"`
}
