package main

import "fmt"

func rechner() {
	fmt.Println("=== Aufgabe 2 Rechner ===")

	a := 10.0
	b := 20.0

	fmt.Printf("Rechnung mit a=%.1f und b=%.1f\n", a, b)
	fmt.Printf("Addition: %.2f\n", a+b)
	fmt.Printf("Subtraktion: %.2f\n", a-b)
	fmt.Printf("Multiplikation: %.2f\n", a*b)

	if b != 0 {
		fmt.Printf("Division: %.2f\n", a/b)
	} else {
		fmt.Println("Division: Fehler - Division durch Null!")
	}
}
