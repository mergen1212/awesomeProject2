package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		// Обработка случая, когда строка короче 100 символов.
		justString = v
	}
}

func main() {
	someFunc()
}
