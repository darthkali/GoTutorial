package main

import "fmt"

func name() {
	fmt.Println("=== Aufgabe 1 Name ===")
	name := "Danny"
	age := 35

	fmt.Println("Hallo, ich bin", name, "und", age, "Jahre alt.")
	fmt.Printf("Hallo, ich bin %s und %d Jahre alt.\n", name, age)
}
