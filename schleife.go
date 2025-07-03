package main

import "fmt"

func schleife() {
	fmt.Println("=== Aufgabe 3 Schleifen ===")

	fmt.Println("=== Zahlen von 1 bis 10 ===")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n=== Gerade Zahlen von 1 bis 20 ===")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\n=== Summe von 1 bis 100 ===")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i // Kurzform für sum = sum + i
	}
	fmt.Printf("Die Summe ist: %d\n", sum)

}
