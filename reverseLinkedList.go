package main

type Node struct {
	value int
	next  *Node
}

// у каждого узла есть указатель на следующий узел.
// нам нужно чтобы указатель указывал не на следующий узел, а на предыдущий

// 1 -> 2 -> 3 -> 4 -> nil
// 4 -> 3 -> 2 -> 1 -> nil

func reverseList(head *Node) *Node {
	// 1. пробегаемся по массиву
	// 2. сохраняем current.next в переменную next
	// 3. сохраняем current как prev в переменную prev
	var prev *Node = nil
	current := head

	for current != nil {
		// 1. Сохраняем указатель на следующий узел
		next := current.next

		// 2. Перепривязываем текущий узел, меняем его указатель на следующий на prev
		current.next = prev

		// 3. Передвигаем prev на текущий узел
		prev = current

		// 4. Переходим к следующему узлу
		current = next
	}

	return prev

}

func main() {

}
