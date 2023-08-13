package addtwonumbers_test

import (
	"testing"

	addtwonumbers "github.com/rzldimam28/leet-code/2-add-two-numbers"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {

	// l1 = [2, 4, 6]
	// l2 = [5, 6, 1]

	l1_3 := addtwonumbers.ListNode{Val: 6}
	l1_2 := addtwonumbers.ListNode{Val: 4, Next: &l1_3}
	l1_1 := addtwonumbers.ListNode{Val: 2, Next: &l1_2}
	l2_3 := addtwonumbers.ListNode{Val: 1}
	l2_2 := addtwonumbers.ListNode{Val: 6, Next: &l2_3}
	l2_1 := addtwonumbers.ListNode{Val: 5, Next: &l2_2}

	res := addtwonumbers.AddTwoNumbers(&l1_1, &l2_1)

	assert.Equal(t, 7, res.Val)
	assert.Equal(t, 0, res.Next.Val)
	assert.Equal(t, 8, res.Next.Next.Val)

}