package logic

import (
	"fmt"
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
		url:      "https://github.com/klovercloud-ci-cd/klovercloud-ci-integration-manager.git",
		expected: "klovercloud-ci-cd",
	}
	testdata.actual, _ = trimUrl(testdata.url)
	if !reflect.DeepEqual(testdata.expected, testdata.actual) {
		fmt.Println(testdata.actual)
		assert.ElementsMatch(t, testdata.expected, testdata.actual)
	}
	_, testdata.actual = trimUrl(testdata.url)
	testdata.expected = "klovercloud-ci-integration-manager"
	if !reflect.DeepEqual(testdata.expected, testdata.actual) {
		fmt.Println(testdata.actual)
		assert.ElementsMatch(t, testdata.expected, testdata.actual)
	}
}
