package main

//Desafio para exibir todos os numeros até 100, que sejam divisiveis por 3

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)

		}

	}
}
