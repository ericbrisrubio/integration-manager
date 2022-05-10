package v1

import "time"

// BitBucketBranches is a list of branches
type BitBucketBranches struct {
	Values []struct {
		Name string `json:"name"`
	} `json:"values"`
}

type BitBucketCommits struct {
	Pagelen int `json:"pagelen"`
	Values  []struct {
		Rendered struct {
			Message struct {
				Raw    string `json:"raw"`
				Markup string `json:"markup"`
				HTML   string `json:"html"`
				Type   string `json:"type"`
			} `json:"message"`
		} `json:"rendered"`
		Hash       string `json:"hash"`
		Repository struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type     string `json:"type"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			UUID     string `json:"uuid"`
		} `json:"repository"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			Patch struct {
				Href string `json:"href"`
			} `json:"patch"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Diff struct {
				Href string `json:"href"`
			} `json:"diff"`
			Approve struct {
				Href string `json:"href"`
			} `json:"approve"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"links"`
		Author struct {
			Raw  string `json:"raw"`
			Type string `json:"type"`
			User struct {
				DisplayName string `json:"display_name"`
				UUID        string `json:"uuid"`
				Links       struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Type      string `json:"type"`
				Nickname  string `json:"nickname"`
				AccountID string `json:"account_id"`
			} `json:"user"`
		} `json:"author,omitempty"`
		Summary struct {
			Raw    string `json:"raw"`
			Markup string `json:"markup"`
			HTML   string `json:"html"`
			Type   string `json:"type"`
		} `json:"summary"`
		Parents []struct {
			Hash  string `json:"hash"`
			Type  string `json:"type"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"parents"`
		Date    time.Time `json:"date"`
		Message string    `json:"message"`
		Type    string    `json:"type"`
	} `json:"values"`
	Next string `json:"next"`
}

// BitbucketRepository contains bitbucket repository information
type BitbucketRepository struct {
	Scm     string      `json:"scm"`
	Website interface{} `json:"website"`
	HasWiki bool        `json:"has_wiki"`
	UUID    string      `json:"uuid"`
	Links   struct {
		Watchers struct {
			Href string `json:"href"`
		} `json:"watchers"`
		Branches struct {
			Href string `json:"href"`
		} `json:"branches"`
		Tags struct {
			Href string `json:"href"`
		} `json:"tags"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Clone []struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"clone"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Source struct {
			Href string `json:"href"`
		} `json:"source"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
		Hooks struct {
			Href string `json:"href"`
		} `json:"hooks"`
		Forks struct {
			Href string `json:"href"`
		} `json:"forks"`
		Downloads struct {
			Href string `json:"href"`
		} `json:"downloads"`
		Pullrequests struct {
			Href string `json:"href"`
		} `json:"pullrequests"`
	} `json:"links"`
	ForkPolicy string `json:"fork_policy"`
	FullName   string `json:"full_name"`
	Name       string `json:"name"`
	Project    struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Type string `json:"type"`
		Name string `json:"name"`
		Key  string `json:"key"`
		UUID string `json:"uuid"`
	} `json:"project"`
	Language   string    `json:"language"`
	CreatedOn  time.Time `json:"created_on"`
	Mainbranch struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"mainbranch"`
	Workspace struct {
		Slug  string `json:"slug"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		UUID string `json:"uuid"`
	} `json:"workspace"`
	HasIssues bool `json:"has_issues"`
	Owner     struct {
		DisplayName string `json:"display_name"`
		UUID        string `json:"uuid"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Type      string `json:"type"`
		Nickname  string `json:"nickname"`
		AccountID string `json:"account_id"`
	} `json:"owner"`
	UpdatedOn   time.Time `json:"updated_on"`
	Size        int       `json:"size"`
	Type        string    `json:"type"`
	Slug        string    `json:"slug"`
	IsPrivate   bool      `json:"is_private"`
	Description string    `json:"description"`
}

// BitbucketWebHookEvent contains github web hook event data
type BitbucketWebHookEvent struct {
	Push struct {
		Changes []struct {
			Forced bool `json:"forced"`
			Old    struct {
				Name  string `json:"name"`
				Links struct {
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
				DefaultMergeStrategy string   `json:"default_merge_strategy"`
				MergeStrategies      []string `json:"merge_strategies"`
				Type                 string   `json:"type"`
				Target               struct {
					Rendered struct {
					} `json:"rendered"`
					Hash  string `json:"hash"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
					Author struct {
						Raw  string `json:"raw"`
						Type string `json:"type"`
						User struct {
							DisplayName string `json:"display_name"`
							UUID        string `json:"uuid"`
							Links       struct {
								Self struct {
									Href string `json:"href"`
								} `json:"self"`
								HTML struct {
									Href string `json:"href"`
								} `json:"html"`
								Avatar struct {
									Href string `json:"href"`
								} `json:"avatar"`
							} `json:"links"`
							Type      string `json:"type"`
							Nickname  string `json:"nickname"`
							AccountID string `json:"account_id"`
						} `json:"user"`
					} `json:"author"`
					Summary struct {
						Raw    string `json:"raw"`
						Markup string `json:"markup"`
						HTML   string `json:"html"`
						Type   string `json:"type"`
					} `json:"summary"`
					Parents []struct {
						Hash  string `json:"hash"`
						Type  string `json:"type"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Date       time.Time `json:"date"`
					Message    string    `json:"message"`
					Type       string    `json:"type"`
					Properties struct {
					} `json:"properties"`
				} `json:"target"`
			} `json:"old"`
			Links struct {
				Commits struct {
					Href string `json:"href"`
				} `json:"commits"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Diff struct {
					Href string `json:"href"`
				} `json:"diff"`
			} `json:"links"`
			Created bool `json:"created"`
			Commits []struct {
				Rendered struct {
				} `json:"rendered"`
				Hash  string `json:"hash"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Comments struct {
						Href string `json:"href"`
					} `json:"comments"`
					Patch struct {
						Href string `json:"href"`
					} `json:"patch"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Diff struct {
						Href string `json:"href"`
					} `json:"diff"`
					Approve struct {
						Href string `json:"href"`
					} `json:"approve"`
					Statuses struct {
						Href string `json:"href"`
					} `json:"statuses"`
				} `json:"links"`
				Author struct {
					Raw  string `json:"raw"`
					Type string `json:"type"`
					User struct {
						DisplayName string `json:"display_name"`
						UUID        string `json:"uuid"`
						Links       struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
							Avatar struct {
								Href string `json:"href"`
							} `json:"avatar"`
						} `json:"links"`
						Type      string `json:"type"`
						Nickname  string `json:"nickname"`
						AccountID string `json:"account_id"`
					} `json:"user"`
				} `json:"author"`
				Summary struct {
					Raw    string `json:"raw"`
					Markup string `json:"markup"`
					HTML   string `json:"html"`
					Type   string `json:"type"`
				} `json:"summary"`
				Parents []struct {
					Hash  string `json:"hash"`
					Type  string `json:"type"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"parents"`
				Date       time.Time `json:"date"`
				Message    string    `json:"message"`
				Type       string    `json:"type"`
				Properties struct {
				} `json:"properties"`
			} `json:"commits"`
			Truncated bool `json:"truncated"`
			Closed    bool `json:"closed"`
			New       struct {
				Name  string `json:"name"`
				Links struct {
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
				DefaultMergeStrategy string   `json:"default_merge_strategy"`
				MergeStrategies      []string `json:"merge_strategies"`
				Type                 string   `json:"type"`
				Target               struct {
					Rendered struct {
					} `json:"rendered"`
					Hash  string `json:"hash"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
					Author struct {
						Raw  string `json:"raw"`
						Type string `json:"type"`
						User struct {
							DisplayName string `json:"display_name"`
							UUID        string `json:"uuid"`
							Links       struct {
								Self struct {
									Href string `json:"href"`
								} `json:"self"`
								HTML struct {
									Href string `json:"href"`
								} `json:"html"`
								Avatar struct {
									Href string `json:"href"`
								} `json:"avatar"`
							} `json:"links"`
							Type      string `json:"type"`
							Nickname  string `json:"nickname"`
							AccountID string `json:"account_id"`
						} `json:"user"`
					} `json:"author"`
					Summary struct {
						Raw    string `json:"raw"`
						Markup string `json:"markup"`
						HTML   string `json:"html"`
						Type   string `json:"type"`
					} `json:"summary"`
					Parents []struct {
						Hash  string `json:"hash"`
						Type  string `json:"type"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Date       time.Time `json:"date"`
					Message    string    `json:"message"`
					Type       string    `json:"type"`
					Properties struct {
					} `json:"properties"`
				} `json:"target"`
			} `json:"new"`
		} `json:"changes"`
	} `json:"push"`
	Actor struct {
		DisplayName string `json:"display_name"`
		UUID        string `json:"uuid"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Type      string `json:"type"`
		Nickname  string `json:"nickname"`
		AccountID string `json:"account_id"`
	} `json:"actor"`
	Repository struct {
		Scm     string      `json:"scm"`
		Website interface{} `json:"website"`
		UUID    string      `json:"uuid"`
		Links   struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Project struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type string `json:"type"`
			Name string `json:"name"`
			Key  string `json:"key"`
			UUID string `json:"uuid"`
		} `json:"project"`
		FullName string `json:"full_name"`
		Owner    struct {
			DisplayName string `json:"display_name"`
			UUID        string `json:"uuid"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type      string `json:"type"`
			Nickname  string `json:"nickname"`
			AccountID string `json:"account_id"`
		} `json:"owner"`
		Workspace struct {
			Slug  string `json:"slug"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			UUID string `json:"uuid"`
		} `json:"workspace"`
		Type      string `json:"type"`
		IsPrivate bool   `json:"is_private"`
		Name      string `json:"name"`
	} `json:"repository"`
}

// BitbucketDirectoryContent contains bitbucket directory data
type BitbucketDirectoryContent struct {
	Pagelen int `json:"pagelen"`
	Values  []struct {
		Path   string `json:"path"`
		Type   string `json:"type"`
		Commit struct {
			Type  string `json:"type"`
			Hash  string `json:"hash"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"commit"`
		Mimetype interface{} `json:"mimetype,omitempty"`
		Links    struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Meta struct {
				Href string `json:"href"`
			} `json:"meta"`
			History struct {
				Href string `json:"href"`
			} `json:"history"`
		} `json:"links,omitempty"`
		EscapedPath string        `json:"escaped_path,omitempty"`
		Attributes  []interface{} `json:"attributes,omitempty"`
		Size        int           `json:"size,omitempty"`
	} `json:"values"`
	Page int `json:"page"`
}

// GetGitDirectoryContent converts BitbucketDirectoryContent object to GitDirectoryContent object
func (directoryContent BitbucketDirectoryContent) GetGitDirectoryContent() GitDirectoryContent {
	gitDirectoryContent := GitDirectoryContent{
		Path:        directoryContent.Values[0].Links.Self.Href,
		Size:        directoryContent.Values[0].Size,
		DownloadURL: directoryContent.Values[0].Links.Self.Href,
	}

	if directoryContent.Values[0].Type == "commit_file" {
		gitDirectoryContent.Type = "file"
	}
	return gitDirectoryContent
}

// BitbucketCreateWebhookRequest contains bitbucket web hook creation data
type BitbucketCreateWebhookRequest struct {
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Active      bool     `json:"active"`
	Events      []string `json:"events"`
}

// BitbucketWebhook contains bitbucket web hook data
type BitbucketWebhook struct {
	ReadOnly    bool   `json:"read_only"`
	Description string `json:"description"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
	URL                  string      `json:"url"`
	CreatedAt            time.Time   `json:"created_at"`
	SkipCertVerification bool        `json:"skip_cert_verification"`
	Source               interface{} `json:"source"`
	HistoryEnabled       bool        `json:"history_enabled"`
	Active               bool        `json:"active"`
	Subject              struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Type     string `json:"type"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		UUID     string `json:"uuid"`
	} `json:"subject"`
	Type   string   `json:"type"`
	Events []string `json:"events"`
	UUID   string   `json:"uuid"`
}

// GetGitWebhook converts BitbucketWebhook object to GitWebhook object
func (webhook BitbucketWebhook) GetGitWebhook() GitWebhook {
	return GitWebhook{
		URL:       webhook.URL,
		CreatedAt: webhook.CreatedAt,
		UpdatedAt: webhook.CreatedAt,
		Active:    webhook.Active,
		Type:      webhook.Type,
		Events:    webhook.Events,
		ID:        webhook.UUID,
	}
}
