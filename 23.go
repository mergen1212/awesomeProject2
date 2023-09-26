package main

import "fmt"

func main() {
	slice := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	fmt.Println("удалить ")
	var i int
	fmt.Scan(&i)
	if i >= 0 && i < len(slice) {

		slice = append(slice[:i], slice[i+1:]...)
		fmt.Println("Срез после удаления:", slice)
	} else {
		fmt.Println("Неверный индекс для удаления")
	}
}
