package v1

import "time"

// GitWebhook contains github web hook data
type GitWebhook struct {
	Type   string   `json:"type"`
	ID     string   `json:"id"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Config struct {
		URL         string `json:"url"`
		InsecureSsl string `json:"insecure_ssl"`
		ContentType string `json:"content_type"`
	} `json:"config"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	URL           string    `json:"url"`
	TestURL       string    `json:"test_url"`
	PingURL       string    `json:"ping_url"`
	DeliveriesURL string    `json:"deliveries_url"`
}

// GitBranches contains github branches
type GitBranches []struct {
	Name string `json:"name"`
}

// DirectoryContentCreatePayload contains directory content create payload
type DirectoryContentCreatePayload struct {
	Message string `json:"message"`
	Content string `json:"content"`
}

// DirectoryContentCreateAndUpdateResponse contains directory content create and update response
type DirectoryContentCreateAndUpdateResponse struct {
	Content struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Sha         string `json:"sha"`
		Size        int    `json:"size"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		GitURL      string `json:"git_url"`
		DownloadURL string `json:"download_url"`
		Type        string `json:"type"`
		Links       struct {
			Self string `json:"self"`
			Git  string `json:"git"`
			HTML string `json:"html"`
		} `json:"_links"`
	} `json:"content"`
	Commit struct {
		Sha     string `json:"sha"`
		NodeID  string `json:"node_id"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
		Author  struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Tree struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		Message string `json:"message"`
		Parents []struct {
			Sha     string `json:"sha"`
			URL     string `json:"url"`
			HTMLURL string `json:"html_url"`
		} `json:"parents"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
}

// DirectoryContentUpdatePayload contains directory content update payload
type DirectoryContentUpdatePayload struct {
	Message string `json:"message"`
	Content string `json:"content"`
	Sha     string `json:"sha"`
}

// GitContent contains github content
type GitContent struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Sha      string `json:"sha"`
	Size     int    `json:"size"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
	Links    struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

// GitDirectoryContent contains github directory data
type GitDirectoryContent struct {
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

type Commit struct {
	URL     string `json:"url"`
	Sha     string `json:"sha"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`
	Commit  struct {
		Message string `json:"message"`
		Author  struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

// GitCommit contains github commit data
type GitCommit struct {
	URL     string `json:"url"`
	Sha     string `json:"sha"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`
	Commit  struct {
		Message string `json:"message"`
		Author  struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}
