package problem0035

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}
type para struct {
	paraOne []int
	paraTwo int
}
type ans struct {
	ansOne int
}

func TestSearchInsert(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{para{[]int{1, 3, 5, 6}, 5}, ans{2}},
		{para{[]int{1, 3, 5, 6}, 2}, ans{1}},
		{para{[]int{1, 3, 5, 6}, 7}, ans{4}},
		{para{[]int{1, 3, 5, 6}, 1}, ans{0}},
	}
	for _, q := range qs {
		ast.Equal(q.ansOne, searchInsert(q.paraOne, q.paraTwo), q.para)
		fmt.Println(q.para)
		fmt.Println("")
	}
}
