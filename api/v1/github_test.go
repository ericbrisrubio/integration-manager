package v1

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_branchExists(t *testing.T) {
	type TestCase struct {
		expected bool
		actual   bool
	}

	var testCases [] TestCase
	var step  v1.Step
	step=v1.Step{Type:enums.BUILD , Params: map[enums.PARAMS]string{"revision":"master"}}

	testCases= append(testCases,TestCase{
		expected: true,
		actual:   branchExists([]v1.Step{step},"zeromsi/branch/master"),
	} )

	step=v1.Step{Type:enums.BUILD , Params: map[enums.PARAMS]string{"revision":"dev"}}

	testCases= append(testCases,TestCase{
		expected: true,
		actual:   branchExists([]v1.Step{step},"zeromsi/branch/dev"),
	} )

	testCases= append(testCases,TestCase{
		expected: false,
		actual:   branchExists([]v1.Step{step},"zeromsi/branch/prod"),
	} )

	for i := 0; i < len(testCases); i++ {
		if !reflect.DeepEqual(testCases[i].expected, testCases[i].actual) {
			assert.ElementsMatch(t, testCases[i].expected, testCases[i].actual)
		}
	}
}
