package v1

// GithubDirectoryContent contains github directory data
type GithubDirectoryContent struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Sha         string      `json:"sha"`
	Size        int         `json:"size"`
	URL         string      `json:"url"`
	HTMLURL     string      `json:"html_url"`
	GitURL      string      `json:"git_url"`
	DownloadURL interface{} `json:"download_url"`
	Type        string      `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}
