package v1

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func InitBitBucketDirectoryConetent(str string) BitbucketDirectoryContent {
	res := BitbucketDirectoryContent{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return BitbucketDirectoryContent{}
	}
	return res
}

// Test function for testing Bitbucket directory content to Git directory content serialization.
func TestBitbucketDirectoryContent_GetGitDirectoryContent(t *testing.T) {
	type TestCase struct {
		data     BitbucketDirectoryContent
		expected []GitDirectoryContent
		actual   []GitDirectoryContent
	}

	var testdata []TestCase

	expec := []GitDirectoryContent{
		{
			Path:        "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/configs/",
			Size:        0,
			DownloadURL: "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/configs/",
			Type:        "",
		},
		{
			Path:        "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml",
			Size:        1114,
			DownloadURL: "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml",
			Type:        "file",
		},
	}
	jsonData := []string{
		`{
		   "pagelen":100,
		   "values":[
			  {
				 "path":"klovercloud/pipeline/configs",
				 "type":"commit_directory",
				 "links":{
					"self":{
					   "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/configs/"
					},
					"meta":{
					   "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/configs/?format=meta"
					}
				 },
				 "commit":{
					"type":"commit",
					"hash":"920d562bac358658e431b66d749dfc6ff74d35dc",
					"links":{
					   "self":{
						  "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/commit/920d562bac358658e431b66d749dfc6ff74d35dc"
					   },
					   "html":{
						  "href":"https://bitbucket.org/shabrulislam2451/testapp/commits/920d562bac358658e431b66d749dfc6ff74d35dc"
					   }
					}
				 }
			  },
			  {
				 "mimetype":null,
				 "links":{
					"self":{
					   "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml"
					},
					"meta":{
					   "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml?format=meta"
					},
					"history":{
					   "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/filehistory/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml"
					}
				 },
				 "escaped_path":"klovercloud/pipeline/pipeline.yaml",
				 "path":"klovercloud/pipeline/pipeline.yaml",
				 "commit":{
					"type":"commit",
					"hash":"920d562bac358658e431b66d749dfc6ff74d35dc",
					"links":{
					   "self":{
						  "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/commit/920d562bac358658e431b66d749dfc6ff74d35dc"
					   },
					   "html":{
						  "href":"https://bitbucket.org/shabrulislam2451/testapp/commits/920d562bac358658e431b66d749dfc6ff74d35dc"
					   }
					}
				 },
				 "attributes":[
					
				 ],
				 "type":"commit_file",
				 "size":1114
			  }
		   ],
		   "page":1
		}`,
	}

	for i := 0; i < len(jsonData); i++ {
		testcase := TestCase{
			data:     InitBitBucketDirectoryConetent(jsonData[i]),
			expected: expec,
		}
		testdata = append(testdata, testcase)
	}

	var gitDirectoryContents []GitDirectoryContent
	for i := 0; i < len(jsonData); i++ {
		for _, each := range testdata[i].data.Values {
			bitbucketDirectoryContent := BitbucketDirectoryContent{}
			bitbucketDirectoryContent.Values = append(bitbucketDirectoryContent.Values, each)
			gitDirectoryContents = append(gitDirectoryContents, bitbucketDirectoryContent.GetGitDirectoryContent())
		}
		testdata[i].actual = gitDirectoryContents
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

func InitBitBucketWebhook(str string) BitbucketWebhook {
	res := BitbucketWebhook{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return BitbucketWebhook{}
	}
	return res
}

// Test function for testing Bitbucket webhook content to Git webhook content serialization.
func TestBitbucketWebhook_GetGitWebhook(t *testing.T) {
	type TestCase struct {
		data     BitbucketWebhook
		expected GitWebhook
		actual   GitWebhook
	}

	var testdata []TestCase

	jsonData := []string{
		`{
		   "read_only":false,
		   "description":"Bitbucket webhook",
		   "links":{
			  "self":{
				 "href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/hooks/%7B9c60218c-75e9-4a1c-b308-338aa68b7dbc%7D"
			  }
		   },
		   "url":"http://1647-103-55-145-83.ngrok.io/api/v1/bitbuckets",
		   "created_at":"2021-12-26T05:11:10.935077Z",
		   "skip_cert_verification":false,
		   "source":null,
		   "history_enabled":false,
		   "active":true,
		   "subject":{
			  "links":{
				 "self":{
					"href":"https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp"
				 },
				 "html":{
					"href":"https://bitbucket.org/shabrulislam2451/testapp"
				 },
				 "avatar":{
					"href":"https://bytebucket.org/ravatar/%7B0e2e8861-6d72-44ee-8ea8-4aabfa0006e5%7D?ts=default"
				 }
			  },
			  "type":"repository",
			  "name":"testapp",
			  "full_name":"shabrulislam2451/testapp",
			  "uuid":"{0e2e8861-6d72-44ee-8ea8-4aabfa0006e5}"
		   },
		   "type":"webhook_subscription",
		   "events":[
			  "repo:push"
		   ],
		   "uuid":"{9c60218c-75e9-4a1c-b308-338aa68b7dbc}"
		}`,
	}

	createdAtStrings := []string{"2021-12-26T05:11:10.935077Z"}
	createdAt := []time.Time{}
	for i := 0; i < len(createdAtStrings); i++ {
		parsedCreatedAt, _ := time.Parse(time.RFC3339, createdAtStrings[i])
		createdAt = append(createdAt, parsedCreatedAt)
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

	for i := 0; i < len(jsonData); i++ {
		testcase := TestCase{
			data:     InitBitBucketWebhook(jsonData[i]),
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
