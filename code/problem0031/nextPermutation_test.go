package problem0031

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	res := []int{4, 3, 2, 1}
	//res:=[]int{1,5,8,4,7,6,5,3,1}
	nextPermutation(res)
	as := assert.New(t)

	as.Equal([]int{1, 5, 8, 5, 1, 3, 4, 6, 7}, res)
}
