package main

type testNode struct {
	value int
	next  *testNode
}

// 1 -> 2 -> 3 -> 4 -> nil
// 4 -> 3 -> 2 -> 1 -> nil

func reverseList(node *testNode) *testNode {
	var prev *testNode = nil

	current := node

	for current != nil {

		next := current.next
		current.next = prev
		prev = current
		current = next

	}

	return prev

}

func main() {

}
