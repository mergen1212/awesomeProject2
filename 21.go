package main

import "fmt"

type LegacyPrinter struct{}

func (lp *LegacyPrinter) Print(s string) {
	fmt.Println("Legacy Printer:", s)
}

type Printer interface {
	PrintMessage(string)
}

type MyPrinter struct{}

func (mp *MyPrinter) PrintMessage(msg string) {
	fmt.Println("My Printer:", msg)
}

type LegacyPrinterAdapter struct {
	LegacyPrinter *LegacyPrinter
}

func (a *LegacyPrinterAdapter) PrintMessage(msg string) {
	a.LegacyPrinter.Print(msg)
}
func main() {
	myPrinter := &MyPrinter{}
	legacyPrinter := &LegacyPrinter{}

	adapter := &LegacyPrinterAdapter{
		LegacyPrinter: legacyPrinter,
	}

	myPrinter.PrintMessage("Hello from My Printer")
	adapter.PrintMessage("Hello from Legacy Printer through Adapter")
}
