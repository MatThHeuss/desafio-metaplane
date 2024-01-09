package listoperations_test

import (
	"github.com/MatThHeuss/desafio-metaplane/listOperations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNode(t *testing.T) {

	t.Run("Should convert a slice to ListNode", func(t *testing.T) {
		// Arrange
		slice := []int{1, 2, 3, 4, 5}

		// Act
		listNode := listoperations.SliceToListNode(slice)
		currentNode := listNode

		// Assert
		for _, val := range slice {
			assert.NotNil(t, currentNode)
			assert.Equal(t, val, currentNode.Val)
			currentNode = currentNode.Next
		}
		assert.Nil(t, currentNode)
	})

	t.Run("Should return nil, list length equals to 0", func(t *testing.T) {
		// Arrange
		slice := []int{}

		// Act
		listNode := listoperations.SliceToListNode(slice)

		// Assert
		assert.Empty(t, listNode)

	})

	t.Run("Should merge two non-empty lists", func(t *testing.T) {
		// Arrange
		l1 := listoperations.SliceToListNode([]int{1, 3, 5})
		l2 := listoperations.SliceToListNode([]int{2, 4, 6})
		expectedValues := []int{1, 2, 3, 4, 5, 6}

		// Act
		mergedList := listoperations.MergeTwoLists(l1, l2)

		// Assert
		for _, val := range expectedValues {
			assert.NotNil(t, mergedList)
			assert.Equal(t, val, mergedList.Val)
			mergedList = mergedList.Next
		}
		assert.Nil(t, mergedList, "Merged list should end after the last expected value")
	})

	t.Run("Should return nil for two empty lists", func(t *testing.T) {
		// Arrange
		l1 := listoperations.SliceToListNode([]int{})
		l2 := listoperations.SliceToListNode([]int{})

		// Act
		mergedList := listoperations.MergeTwoLists(l1, l2)

		// Assert
		assert.Nil(t, mergedList)
	})

	t.Run("Should merge two lists with different lengths", func(t *testing.T) {
		// Arrange
		l1 := listoperations.SliceToListNode([]int{1, 3, 5, 7})
		l2 := listoperations.SliceToListNode([]int{2, 4})
		expectedValues := []int{1, 2, 3, 4, 5, 7}

		// Act
		mergedList := listoperations.MergeTwoLists(l1, l2)

		// Assert
		for _, val := range expectedValues {
			assert.NotNil(t, mergedList)
			assert.Equal(t, val, mergedList.Val)
			mergedList = mergedList.Next
		}
		assert.Nil(t, mergedList, "Merged list should end after the last expected value")
	})

}
