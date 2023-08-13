package two_sum_test

import (
	"testing"

	two_sum "github.com/rzldimam28/leet-code/1-two-sum"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	testCases := []struct{
		Name string
		Nums []int
		Target int
		Expected []int
	}{
		{
			Name: "Test Cases I",
			Nums: []int{2, 7, 11, 15},
			Target: 9,
			Expected: []int{0, 1},
		},
		{
			Name: "Test Cases II",
			Nums: []int{3, 2, 4},
			Target: 6,
			Expected: []int{1, 2},
		},
		{
			Name: "Test Cases I",
			Nums: []int{3, 3},
			Target: 6,
			Expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := two_sum.TwoSum(tc.Nums, tc.Target)

			assert.Equal(t, tc.Expected, actual)
		})
	}
	
}