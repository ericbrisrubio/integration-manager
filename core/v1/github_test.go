package v1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

// Test function for testing Github directory content to Git directory content serialization.
func TestGithubDirectoryContent_GetGitDirectoryContent(t *testing.T) {
	type TestCase struct {
		data     GithubDirectoryContent
		expected GitDirectoryContent
		actual   GitDirectoryContent
	}

	var testdata []TestCase

	name := []string{"configs", "pipeline.yaml"}
	path := []string{"klovercloud/pipeline/configs", "klovercloud/pipeline/pipeline.yaml"}
	sha := []string{"d5391ad07f83adfb1fbba0463527484d11d5b9f4", "4e9658a477df11cf6619305b5a1ac897e1b2e42f"}
	size := []int{0, 1122}
	url := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9", "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9"}
	htmlUrl := []string{"https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs", "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml"}
	gitUrl := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4", "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f"}
	downloadUrl := []interface{}{nil, "https://raw.githubusercontent.com/flameOfDimitry/TestApp/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml"}
	typeData := []string{"dir", "file"}
	// Links struct data
	self := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9", "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9"}
	git := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4", "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f"}
	html := []string{"https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs", "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml"}

	type Links struct {
		Self string
		Git  string
		HTML string
	}

	expec := []GitDirectoryContent{
		{Name: "configs",
			Path:        "klovercloud/pipeline/configs",
			Sha:         "d5391ad07f83adfb1fbba0463527484d11d5b9f4",
			Size:        0,
			URL:         "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9",
			HTMLURL:     "https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs",
			GitURL:      "https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4",
			DownloadURL: nil,
			Type:        "dir",
			Links: struct {
				Self string `json:"self"`
				Git  string `json:"git"`
				HTML string `json:"html"`
			}(Links{
				Self: "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9",
				Git:  "https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4",
				HTML: "https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs",
			}),
		},
		{Name: "pipeline.yaml",
			Path:        "klovercloud/pipeline/pipeline.yaml",
			Sha:         "4e9658a477df11cf6619305b5a1ac897e1b2e42f",
			Size:        1122,
			URL:         "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9",
			HTMLURL:     "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
			GitURL:      "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f",
			DownloadURL: "https://raw.githubusercontent.com/flameOfDimitry/TestApp/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
			Type:        "file",
			Links: struct {
				Self string `json:"self"`
				Git  string `json:"git"`
				HTML string `json:"html"`
			}(Links{
				Self: "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9",
				Git:  "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f",
				HTML: "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
			}),
		},
	}

	for i := 0; i < len(name); i++ {
		testcase := TestCase{
			data: GithubDirectoryContent{
				Name:        name[i],
				Path:        path[i],
				Sha:         sha[i],
				Size:        size[i],
				URL:         url[i],
				HTMLURL:     htmlUrl[i],
				GitURL:      gitUrl[i],
				DownloadURL: downloadUrl[i],
				Type:        typeData[i],
				Links: struct {
					Self string `json:"self"`
					Git  string `json:"git"`
					HTML string `json:"html"`
				}(Links{
					Self: self[i],
					Git:  git[i],
					HTML: html[i],
				}),
			},
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(name); i++ {
		testdata[i].actual = testdata[i].data.GetGitDirectoryContent()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

// Test function for testing Github webhook content to Git webhook content serialization.
func TestGithubWebhook_GetGitWebhook(t *testing.T) {
	type TestCase struct {
		data     GithubWebhook
		expected GitWebhook
		actual   GitWebhook
	}

	var testdata []TestCase

	typeData := []string{"Repository"}
	id := []int{334715711}
	active := []bool{true}
	events := [][]string{{"delete", "push", "release"}}

	// Config struct data
	config_url := []string{"http://2756-103-55-145-88.ngrok.io/api/v1/githubs"}
	config_insecureSsl := []string{"0"}
	config_contentType := []string{"form"}

	parsedUpdatedTime, _ := time.Parse(time.RFC3339, "2021-12-23T10:13:30Z")
	updatedAt := []time.Time{parsedUpdatedTime}
	parsedCreatedTime, _ := time.Parse(time.RFC3339, "2021-12-23T10:13:30Z")
	createdAt := []time.Time{parsedCreatedTime}

	url := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711"}
	testUrl := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/test"}
	pingUrl := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/pings"}
	deliveriesURL := []string{"https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/deliveries"}

	type Config struct {
		URL         string
		InsecureSsl string
		ContentType string
	}

	expec := []GitWebhook{
		{Type: "Repository",
			ID:     "334715711",
			Active: true,
			Events: []string{"delete", "push", "release"},
			Config: struct {
				URL         string `json:"url"`
				InsecureSsl string `json:"insecure_ssl"`
				ContentType string `json:"content_type"`
			}(Config{
				URL:         "http://2756-103-55-145-88.ngrok.io/api/v1/githubs",
				InsecureSsl: "0",
				ContentType: "form",
			}),
			UpdatedAt:     updatedAt[0],
			CreatedAt:     createdAt[0],
			URL:           "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711",
			TestURL:       "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/test",
			PingURL:       "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/pings",
			DeliveriesURL: "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/deliveries",
		},
	}

	for i := 0; i < len(id); i++ {
		testcase := TestCase{
			data: GithubWebhook{
				Type:   typeData[i],
				ID:     id[i],
				Active: active[i],
				Events: events[i],
				Config: struct {
					URL         string `json:"url"`
					InsecureSsl string `json:"insecure_ssl"`
					ContentType string `json:"content_type"`
				}(Config{
					URL:         config_url[i],
					InsecureSsl: config_insecureSsl[i],
					ContentType: config_contentType[i],
				}),
				UpdatedAt:     updatedAt[i],
				CreatedAt:     createdAt[i],
				URL:           url[i],
				TestURL:       testUrl[i],
				PingURL:       pingUrl[i],
				DeliveriesURL: deliveriesURL[i],
			},
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(id); i++ {
		testdata[i].actual = testdata[i].data.GetGitWebhook()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
