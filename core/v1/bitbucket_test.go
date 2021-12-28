package v1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

// Test function for testing Bitbucket directory content to Git directory content serialization.
func TestBitbucketDirectoryContent_GetGitDirectoryContent(t *testing.T) {
	type TestCase struct {
		data     BitbucketDirectoryContent
		expected GitDirectoryContent
		actual   GitDirectoryContent
	}

	var testdata []TestCase

	pageLen := []int{100}

	// Values struct data
	path := []string{"klovercloud/pipeline/configs"}
	typeData := []string{"commit_directory"}
	// Commit struct data
	commitType := []string{"commit"}
	commitHash := []string{"920d562bac358658e431b66d749dfc6ff74d35dc"}
	// Commit Links struct data
	// Self struct data
	commitLinksSelfHref := []string{"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/commit/920d562bac358658e431b66d749dfc6ff74d35dc"}
	// HTML struct data
	commitLinksHtmlHref := []string{"https://bitbucket.org/shabrulislam2451/testapp/commits/920d562bac358658e431b66d749dfc6ff74d35dc"}

	mimetype := []interface{}{nil}

	// Links struct data
	// Self struct data
	linksSelfHref := []string{"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml"}
	// Meta struct data
	linksMetaHref := []string{"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml?format=meta"}
	// History struct data
	linksHistoryHref := []string{"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/filehistory/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml"}

	escapedPath := []string{"klovercloud/pipeline/pipeline.yaml"}
	attributes := [][]interface{}{{}}
	size := []int{1114}

	page := []int{1}

	type Self struct {
		Href string `json:"href"`
	}
	type HTML struct {
		Href string `json:"href"`
	}
	type Meta struct {
		Href string `json:"href"`
	}
	type History struct {
		Href string `json:"href"`
	}
	type commitLinks struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	}
	type Commit struct {
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
	}
	type Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Meta struct {
			Href string `json:"href"`
		} `json:"meta"`
		History struct {
			Href string `json:"href"`
		} `json:"history"`
	}
	type Values []struct {
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
	}

	valueType := []string{}
	for i := 0; i < len(typeData); i++ {
		if typeData[i] == "commit_file" {
			valueType = append(valueType, "file")
		} else {
			valueType = append(valueType, "")
		}
	}
	expec := []GitDirectoryContent{
		{
			Path:        "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml",
			Size:        1114,
			DownloadURL: "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml",
			Type:        valueType[0],
		},
	}

	for i := 0; i < len(pageLen); i++ {
		testcase := TestCase{
			data: BitbucketDirectoryContent{
				Pagelen: pageLen[i],
				Values: Values{
					{
						Path: path[i],
						Type: valueType[i],
						Commit: Commit{
							Type: commitType[i],
							Hash: commitHash[i],
							Links: commitLinks{
								Self: Self{Href: commitLinksSelfHref[i]},
								HTML: HTML{Href: commitLinksHtmlHref[i]},
							},
						},
						Mimetype: mimetype[i],
						Links: Links{
							Self:    Self{Href: linksSelfHref[i]},
							Meta:    Meta{Href: linksMetaHref[i]},
							History: History{Href: linksHistoryHref[i]},
						},
						EscapedPath: escapedPath[i],
						Attributes:  attributes[i],
						Size:        size[i],
					},
				},
				Page: page[i],
			},
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(pageLen); i++ {
		testdata[i].actual = testdata[i].data.GetGitDirectoryContent()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

// Test function for testing Bitbucket webhook content to Git webhook content serialization.
func TestBitbucketWebhook_GetGitWebhook(t *testing.T) {
	type TestCase struct {
		data     BitbucketWebhook
		expected GitWebhook
		actual   GitWebhook
	}

	var testdata []TestCase

	readOnly := []bool{false}
	description := []string{"Bitbucket webhook"}

	// Links struct data
	// Self struct data
	linksSelfHref := []string{"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/hooks/%7B9c60218c-75e9-4a1c-b308-338aa68b7dbc%7D"}

	url := []string{"http://1647-103-55-145-83.ngrok.io/api/v1/bitbuckets"}
	parsedCreatedTime, _ := time.Parse(time.RFC3339, "2021-12-26T05:11:10.935077Z")
	createdAt := []time.Time{parsedCreatedTime}
	skipCertifiedVerification := []bool{false}
	source := []interface{}{nil}
	historyEnabled := []bool{false}
	active := []bool{true}

	// Subject struct data
	// Links struct data
	// Self struct data
	subjectLinksSelfHref := []string{"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp"}
	// HTML struct data
	subjectLinksHtmlHref := []string{"https://bitbucket.org/shabrulislam2451/testapp"}
	// Avatar struct data
	subjectLinksAvatarHref := []string{"https://bytebucket.org/ravatar/%7B0e2e8861-6d72-44ee-8ea8-4aabfa0006e5%7D?ts=default"}
	subjectType := []string{"repository"}
	subjectName := []string{"testapp"}
	subjectFullName := []string{"shabrulislam2451/testapp"}
	subjectUuid := []string{"{0e2e8861-6d72-44ee-8ea8-4aabfa0006e5}"}

	typeData := []string{"webhook_subscription"}
	events := [][]string{{"repo:push"}}
	uuid := []string{"{9c60218c-75e9-4a1c-b308-338aa68b7dbc}"}

	type Self struct {
		Href string `json:"href"`
	}
	type HTML struct {
		Href string `json:"href"`
	}
	type Avatar struct {
		Href string `json:"href"`
	}
	type Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	}
	type subjectLinks struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	}
	type Subject struct {
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
	}

	expec := []GitWebhook{
		{
			URL:       "http://1647-103-55-145-83.ngrok.io/api/v1/bitbuckets",
			CreatedAt: createdAt[0],
			UpdatedAt: createdAt[0],
			Active:    true,
			Type:      "webhook_subscription",
			Events:    []string{"repo:push"},
			ID:        "{9c60218c-75e9-4a1c-b308-338aa68b7dbc}",
		},
	}

	for i := 0; i < len(readOnly); i++ {
		testcase := TestCase{
			data: BitbucketWebhook{
				ReadOnly:    readOnly[i],
				Description: description[i],
				Links: Links{
					Self: Self{Href: linksSelfHref[i]},
				},
				URL:                  url[i],
				CreatedAt:            createdAt[i],
				SkipCertVerification: skipCertifiedVerification[i],
				Source:               source[i],
				HistoryEnabled:       historyEnabled[i],
				Active:               active[i],
				Subject: Subject{
					Links: subjectLinks{
						Self:   Self{Href: subjectLinksSelfHref[i]},
						HTML:   HTML{Href: subjectLinksHtmlHref[i]},
						Avatar: Avatar{Href: subjectLinksAvatarHref[i]},
					},
					Type:     subjectType[i],
					Name:     subjectName[i],
					FullName: subjectFullName[i],
					UUID:     subjectUuid[i],
				},
				Type:   typeData[i],
				Events: events[i],
				UUID:   uuid[i],
			},
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(readOnly); i++ {
		testdata[i].actual = testdata[i].data.GetGitWebhook()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
