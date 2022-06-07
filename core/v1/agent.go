package v1

// Agent contains payload data
type Agent struct {
	Name            string `bson:"name" json:"name"`
	ApiVersion      string `bson:"api_version" json:"api_version"`
	TerminalBaseUrl string `bson:"terminal_base_url" json:"terminal_base_url"`
}
