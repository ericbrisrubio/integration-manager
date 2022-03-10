package v1

import (
	"fmt"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestTrimUrl(t *testing.T) {
	type TestCase struct {
		data     string
		expected string
		actual   string
	}

	testCases := []TestCase{
		TestCase{
			data:     "https://bitbucket.org/shabrulislam2451/testapp/src/master/",
			expected: "https://bitbucket.org/shabrulislam2451/testapp",
			actual:   UrlFormatter("https://bitbucket.org/shabrulislam2451/testapp/src/master/"),
		},
		TestCase{
			data:     "https://shabrulislam2451@bitbucket.org/shabrulislam2451/testapp.git",
			expected: "https://bitbucket.org/shabrulislam2451/testapp",
			actual:   UrlFormatter("https://shabrulislam2451@bitbucket.org/shabrulislam2451/testapp.git"),
		},
		TestCase{
			data:     "https://github.com/shabrul2451/test-app.git",
			expected: "https://github.com/shabrul2451/test-app",
			actual:   UrlFormatter("https://github.com/shabrul2451/test-app.git"),
		},
	}
	for i := 0; i < len(testCases); i++ {
		if !reflect.DeepEqual(testCases[i].expected, testCases[i].actual) {
			fmt.Println(testCases[i].actual, testCases[i].expected)
			assert.ElementsMatch(t, testCases[i].expected, testCases[i].actual)
		}
	}
}

func Test_branchExists(t *testing.T) {
	type TestCase struct {
		expected bool
		actual   bool
	}

	var testCases []TestCase
	var step v1.Step
	step = v1.Step{Type: enums.BUILD, Params: map[enums.PARAMS]string{"revision": "master"}}

	testCases = append(testCases, TestCase{
		expected: true,
		actual:   BranchExists([]v1.Step{step}, "zeromsi/branch/master", "GITHUB"),
	})

	step = v1.Step{Type: enums.BUILD, Params: map[enums.PARAMS]string{"revision": "dev"}}

	testCases = append(testCases, TestCase{
		expected: true,
		actual:   BranchExists([]v1.Step{step}, "zeromsi/branch/dev", "GITHUB"),
	})

	testCases = append(testCases, TestCase{
		expected: false,
		actual:   BranchExists([]v1.Step{step}, "zeromsi/branch/prod", "GITHUB"),
	})

	step = v1.Step{Type: enums.BUILD, Params: map[enums.PARAMS]string{"revision": "master"}}

	testCases = append(testCases, TestCase{
		expected: true,
		actual:   BranchExists([]v1.Step{step}, "master", "BIT_BUCKET"),
	})

	step = v1.Step{Type: enums.BUILD, Params: map[enums.PARAMS]string{"revision": "dev"}}

	testCases = append(testCases, TestCase{
		expected: true,
		actual:   BranchExists([]v1.Step{step}, "dev", "BIT_BUCKET"),
	})

	testCases = append(testCases, TestCase{
		expected: false,
		actual:   BranchExists([]v1.Step{step}, "prod", "BIT_BUCKET"),
	})

	for i := 0; i < len(testCases); i++ {
		if !reflect.DeepEqual(testCases[i].expected, testCases[i].actual) {
			assert.ElementsMatch(t, testCases[i].expected, testCases[i].actual)
		}
	}
}
