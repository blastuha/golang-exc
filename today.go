package main

type LinkedNode struct {
	value int
	next  *LinkedNode
}

func addToFront(head **LinkedNode, value int) {
	// Создаем новый узел
	newNode := &LinkedNode{
		value: value,
		next:  *head, // newNode.next сохраняет старую ссылку на голову списка (она не изменится строчкой ниже); здесь мы сохраняем АДРЕС старой головы списка.
	}
	// Затем мы присваиваем newNode в *head. Таким образом, мы изменяем указатель на первый узел так, чтобы он теперь указывал на newNode.
	*head = newNode
	// присваивает переменной linkedHead новое значение — это новый узел newNode.
}

func main() {
	var linkedHead *LinkedNode
	addToFront(&linkedHead, 1)

}
