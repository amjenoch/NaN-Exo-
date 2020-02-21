package main

import "fmt"

// Factoriel en Recursitive
func Factoriel(Nombre int) int{
	
	if Nombre >= -1 && Nombre <= 1 {
		return Nombre
	}
	if Nombre > 1 {
	return Nombre * Factoriel(Nombre - 1)
	}
	if Nombre < -1 {
		return -Nombre * Factoriel(Nombre + 1)
	}
        
	return Nombre
}

func main()  {
	var Nombre int
	fmt.Print("Entrez un nombre:")
	fmt.Scan(&Nombre)
	fmt.Println("Factoriel de",Nombre, "est: ")
	fmt.Println(Factoriel(Nombre))
}


