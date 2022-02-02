package v1

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

// Initialize Github directory content from JSON data
func InitGithubDirectoryContent(str string) GithubDirectoryContent {
	res := GithubDirectoryContent{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return GithubDirectoryContent{}
	}
	return res
}

// Test function for testing Github directory content to Git directory content serialization.
func TestGithubDirectoryContent_GetGitDirectoryContent(t *testing.T) {
	type TestCase struct {
		data     GithubDirectoryContent
		expected GitDirectoryContent
		actual   GitDirectoryContent
	}

	var testdata []TestCase

	jsonData := []string{
		`{
			"name": "configs",
			"path": "klovercloud/pipeline/configs",
			"sha": "d5391ad07f83adfb1fbba0463527484d11d5b9f4",
			"size": 0,
			"url": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9",
			"html_url": "https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs",
			"git_url": "https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4",
			"download_url": null,
			"type": "dir",
			"_links": {
			  "self": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9",
			  "git": "https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4",
			  "html": "https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs"
			}
		  }`,
		`{
			"name": "pipeline.yaml",
			"path": "klovercloud/pipeline/pipeline.yaml",
			"sha": "4e9658a477df11cf6619305b5a1ac897e1b2e42f",
			"size": 1122,
			"url": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9",
			"html_url": "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
			"git_url": "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f",
			"download_url": "https://raw.githubusercontent.com/flameOfDimitry/TestApp/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
			"type": "file",
			"_links": {
			  "self": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9",
			  "git": "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f",
			  "html": "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml"
			}
		  }`,
	}

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

	for i := 0; i < len(jsonData); i++ {
		testcase := TestCase{
			data:     InitGithubDirectoryContent(jsonData[i]),
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(jsonData); i++ {
		testdata[i].actual = testdata[i].data.GetGitDirectoryContent()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

func InitGithubWebhook(str string) GithubWebhook {
	res := GithubWebhook{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return GithubWebhook{}
	}
	return res
}

// Test function for testing Github webhook content to Git webhook content serialization.
func TestGithubWebhook_GetGitWebhook(t *testing.T) {
	type TestCase struct {
		data     GithubWebhook
		expected GitWebhook
		actual   GitWebhook
	}

	var testdata []TestCase

	jsonData := []string{
		`{
		  "type": "Repository",
		  "id": 334715711,
		  "name": "web",
		  "active": true,
		  "events": [
			"delete",
			"push",
			"release"
		  ],
		  "config": {
			"url": "http://2756-103-55-145-88.ngrok.io/api/v1/githubs",
			"insecure_ssl": "0",
			"content_type": "form"
		  },
		  "updated_at": "2021-12-23T10:13:30Z",
		  "created_at": "2021-12-23T10:13:30Z",
		  "url": "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711",
		  "test_url": "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/test",
		  "ping_url": "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/pings",
		  "deliveries_url": "https://api.github.com/repos/flameOfDimitry/TestApp/hooks/334715711/deliveries",
		  "last_response": {
			"code": null,
			"status": "unused",
			"message": null
		  }
		}`,
	}

	type Config struct {
		URL         string
		InsecureSsl string
		ContentType string
	}

	updatedAtStrings := []string{"2021-12-23T10:13:30Z"}
	updatedAt := []time.Time{}
	createdAtStrings := []string{"2021-12-23T10:13:30Z"}
	createdAt := []time.Time{}

	for i := 0; i < len(updatedAtStrings); i++ {
		parsedUpdatedTime, _ := time.Parse(time.RFC3339, updatedAtStrings[i])
		updatedAt = append(updatedAt, parsedUpdatedTime)
	}
	for i := 0; i < len(createdAtStrings); i++ {
		parsedCreatedAt, _ := time.Parse(time.RFC3339, createdAtStrings[i])
		createdAt = append(createdAt, parsedCreatedAt)
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

	for i := 0; i < len(jsonData); i++ {
		testcase := TestCase{
			data:     InitGithubWebhook(jsonData[i]),
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i < len(jsonData); i++ {
		testdata[i].actual = testdata[i].data.GetGitWebhook()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
