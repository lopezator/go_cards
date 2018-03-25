package main

import "fmt"

func main() {
	//Aquí el tipo no es necesario porque se "infiere"
	//var card string = "Ace of Spades"
	//var card = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card)
}

/**

Declarar el tipo de la variable explicitamente:

--> var card string = "Ace of Spades"

En este caso, no es necesario, ya que puede ser inferido automáticamente.

--> var card = "Ace of Spades

Es equivalente a:

--> card := "Ace of Spades"

Pero el operador ":=" no puede utilizarse fuera de un bloque "func"

Una vez la variable está ya declarada, no puede utilizarse "var" ni ":=" se utiliza el nombre de
la variable directamente para reasignarla.

--> card = "Five of diamonds"

 */