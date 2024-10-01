package main

import "fmt"

type LinkedNode struct {
	value int
	next  *LinkedNode
}

func addToFront(head **LinkedNode, value int) {
	// Ты разыменовываешь двойной указатель **LinkedNode один раз, чтобы получить саму переменную linkedHead, которая имеет тип *LinkedNode. То есть *head — это содержимое переменной linkedHead, а оно изначально равно nil (список пуст).

	newNode := &LinkedNode{
		value: value,
		next:  *head, // разыменовываем один раз, чтобы получить указатель на linkedHead
	}
	*head = newNode // обновляем указатель на новый узел
}

func addToEnd2(head **LinkedNode, value int) {
	// 1. создаем новый узел
	newNode := &LinkedNode{
		value: value,
		next:  nil,
	}

	// 2. Если список пуст (head указывает на nil), новый узел становится первым элементом
	if *head == nil {
		*head = newNode
		return
	}

	// 3. Ищем последний узел, чей next равен nil
	current := *head
	for current.next != nil {
		current = current.next
	}

	// 4. Добавляем новый узел в конец списка
	current.next = newNode
}

func printList2(head *LinkedNode) {
	current := head

	for current != nil {
		fmt.Println("current", current.value)
		current = current.next
	}
}

func main() {
	var linkedHead *LinkedNode = nil
	addToFront(&linkedHead, 1)
	addToFront(&linkedHead, 2)
	printList2(linkedHead)
}
