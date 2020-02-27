package problem0014

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	ast := assert.New(t)
	questions := []string{
		"flower",
		"flow",
		"flight",
	}
	ans := "fl"
	ast.Equal(ans, longestCommonPrefix1(questions), questions)
}
