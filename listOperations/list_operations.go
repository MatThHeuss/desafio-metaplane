package listoperations

type ListNode struct {
	Val  int
	Next *ListNode
}

// Estrutura para receber listas
type List struct {
	Numbers []int `json:"numbers"`
}

// Função para mesclar duas listas
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

// Função para converter slice em ListNode
func SliceToListNode(numbers []int) *ListNode {
	if len(numbers) == 0 {
		return nil
	}
	head := &ListNode{Val: numbers[0]}
	current := head
	for _, val := range numbers[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}
