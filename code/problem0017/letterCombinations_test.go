package problem0017

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}
type para struct {
	one string
}
type ans struct {
	one []string
}

func TestLetterCombinations(t *testing.T) {
	ast := assert.New(t)
	questions := []question{
		{para{"23"}, ans{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}}},
		{para{""}, ans{nil}},
	}
	for _, q := range questions {
		ast.Equal(q.ans.one, letterCombinations2(q.para.one), q.para.one)
	}
}
func BenchmarkLetter(b *testing.B) {
	questions := []question{
		{para{"23"}, ans{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}}},
		{para{""}, ans{nil}},
	}
	for i := 0; i < b.N; i++ {
		letterCombinations2(questions[0].para.one)
	}

}
