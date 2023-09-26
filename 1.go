package main

import "fmt"

// Создаем структуру Human

type Human struct {
	Name    string
	Age     int
	Address string
}

// Добавляем метод в структуру Human

func (h *Human) Speak() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет, и я живу по адресу %s\n", h.Name, h.Age, h.Address)
}

// Создаем структуру Action и встраиваем в нее структуру Human

type Action struct {
	Human
	Activity string
}

func main() {

	// Создаем объект Action с доступом к полям и методам Human

	action := Action{
		Human:    Human{Name: "Иван", Age: 30, Address: "ул. Ленина, 123"},
		Activity: "гулять в парке",
	}

	// Мы можем вызвать метод Speak() из структуры Human через структуру Action

	action.Speak()

	// Используем поля и методы из структуры Action

	fmt.Printf("%s любит %s\n", action.Name, action.Activity)
}
