package v1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGithub_setAccessModeForBuild(t *testing.T) {
	type TestCase struct {
		data     string
		expected string
		actual   string
	}
	var testCases = []TestCase{
		{
			data:     "ReadWriteOnce",
			expected: "ReadWriteOnce",
		},
		{
			data:     "ReadWriteMany",
			expected: "ReadWriteMany",
		},
		{
			data:     "ReadOnlyMany",
			expected: "ReadOnlyMany",
		},
		{
			data:     "ReadWriteOncePod",
			expected: "ReadWriteOncePod",
		},
		{
			data:     "",
			expected: "ReadWriteOnce",
		},
	}
	for i, _ := range testCases {
		testCases[i].actual = setAccessModeForBuild(testCases[i].data)
		if !reflect.DeepEqual(testCases[i].expected, testCases[i].actual) {
			fmt.Println(testCases[i].actual)
			assert.ElementsMatch(t, testCases[i].expected, testCases[i].actual)
		}
	}
}
