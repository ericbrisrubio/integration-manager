package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_trimUrl(t *testing.T) {
	type TestCase struct {
		url      string
		expected string
		actual   string
	}

	testdata := TestCase{
		url:      "https://github.com/klovercloud-ci-cd/integration-manager.git",
		expected: "klovercloud-ci-cd",
	}
	testdata.actual, _ = v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(testdata.url)
	if !reflect.DeepEqual(testdata.expected, testdata.actual) {
		assert.ElementsMatch(t, testdata.expected, testdata.actual)
	}
	_, testdata.actual = v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(testdata.url)
	testdata.expected = "integration-manager"
	if !reflect.DeepEqual(testdata.expected, testdata.actual) {
		assert.ElementsMatch(t, testdata.expected, testdata.actual)
	}
}
