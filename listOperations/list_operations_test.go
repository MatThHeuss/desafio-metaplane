package listoperations_test

import (
	"github.com/MatThHeuss/desafio-metaplane/listOperations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNode(t *testing.T) {

	t.Run("Should convert a slice to ListNode", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		listNode := listoperations.SliceToListNode(slice)

		currentNode := listNode
		for _, val := range slice {
			assert.NotNil(t, currentNode, "ListNode should not be nil")
			assert.Equal(t, val, currentNode.Val, "ListNode value mismatch")
			currentNode = currentNode.Next
		}

		assert.Nil(t, currentNode, "ListNode should end after the last slice element")
	})

	t.Run("Should return nil, list length equals to 0", func(t *testing.T) {
		slice := []int{}
		listNode := listoperations.SliceToListNode(slice)
		assert.Empty(t, listNode)

	})

	t.Run("Should merge two non-empty lists", func(t *testing.T) {
		l1 := listoperations.SliceToListNode([]int{1, 3, 5})
		l2 := listoperations.SliceToListNode([]int{2, 4, 6})
		mergedList := listoperations.MergeTwoLists(l1, l2)

		expectedValues := []int{1, 2, 3, 4, 5, 6}
		for _, val := range expectedValues {
			assert.NotNil(t, mergedList, "Merged list should not be nil")
			assert.Equal(t, val, mergedList.Val, "Value mismatch in merged list")
			mergedList = mergedList.Next
		}
		assert.Nil(t, mergedList, "Merged list should end after the last expected value")
	})

	t.Run("Should return nil for two empty lists", func(t *testing.T) {
		l1 := listoperations.SliceToListNode([]int{})
		l2 := listoperations.SliceToListNode([]int{})
		mergedList := listoperations.MergeTwoLists(l1, l2)

		assert.Nil(t, mergedList)
	})

	t.Run("Should merge two lists with different lengths", func(t *testing.T) {
		// First list is longer than the second
		l1 := listoperations.SliceToListNode([]int{1, 3, 5, 7})
		l2 := listoperations.SliceToListNode([]int{2, 4})
		mergedList := listoperations.MergeTwoLists(l1, l2)

		// The expected merge result
		expectedValues := []int{1, 2, 3, 4, 5, 7}
		for _, val := range expectedValues {
			assert.NotNil(t, mergedList, "Merged list should not be nil")
			assert.Equal(t, val, mergedList.Val, "Value mismatch in merged list")
			mergedList = mergedList.Next
		}
		assert.Nil(t, mergedList, "Merged list should end after the last expected value")
	})

}
