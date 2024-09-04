package main

import "fmt"

type Notifier interface {
	Notify(message string) string
}

type EmailNotifier struct {
	email string
}

func (n EmailNotifier) Notify(message string) string {
	return fmt.Sprintf("Отправка email на %s %s", n.email, message)
}

func main() {
	emailNotifier := EmailNotifier{email: "shevnin.boris2@yandex.ru"}
	fmt.Println(emailNotifier, emailNotifier.Notify("Соси бибу лох"))
}
