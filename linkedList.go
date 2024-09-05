package main

import "fmt"

// Создайте односвязный список на языке Golang, который будет хранить целые числа. Реализуйте функции для:
//Добавления нового элемента в конец списка.
//Печати всех элементов списка.
//Удаления элемента из списка по значению.

type LinkedList struct {
	value int
	next  *LinkedList
}

func addToStart(number int, head *LinkedList) *LinkedList {
	// 1. принимаем голову в аргументах
	// 2. создаем новый узел, чей next ссылается на эту голову
	// 3. тем самым мы создаем новую голову, которая ссылается на старую голову
	// 4. возвращаем новую голову
	var newNode *LinkedList = &LinkedList{
		value: number,
		next:  head,
	}

	return newNode
}

func addToEnd(number int, head *LinkedList) *LinkedList {
	// 1. принимает голову в аргумента + создаем новый узел
	// 2. так как узел будет конечным, его next = nil
	// 3. если голова = nil, то значит список пуст и мы возвращаем созданный узел
	// 4. пока current.next != nil, мы в цикле присваиваем новое значение current, пока не дойдем до предпоследнего элемента. тем самым current станет последним элементом
	// 5. далее указываем что current.next ссылается newNode

	var newNode *LinkedList = &LinkedList{
		value: number,
		next:  nil,
	}

	// если список пуст
	if head == nil {
		return newNode
	}

	current := head
	// Пока current.next != nil, мы в цикле присваиваем новое значение current, пока не дойдем до последнего элемента

	for current.next != nil {
		// Итерация возможна, так как мы постоянно обновляем current через условие ниже
		current = current.next
	}
	// Теперь current указывает на последний элемент, добавляем новый узел в конец списка
	current.next = newNode

	return newNode
}

func printList(head *LinkedList) {
	current := head

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func getListLength(head *LinkedList) int {
	var n int = 0
	current := head

	for current != nil {
		n += 1
		current = current.next
	}
	fmt.Println("n", n)
	return n
}

func main() {

	var head *LinkedList = nil

	for i := 0; i < 4; i++ {
		head = addToStart(i, head)
	}

	// для теста добавляю в начало списка число 10
	var newHead *LinkedList = addToStart(10, head)

	printList(newHead)
	getListLength(newHead)

}
