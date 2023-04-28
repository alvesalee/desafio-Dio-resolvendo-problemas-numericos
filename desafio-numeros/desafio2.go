package main

import "fmt"

func main() {

	//PRIMEIRA TENTATIVA DEU RUIM  :'(

	//	for i := 0; i <= 100; i++ {
	//		if i%3 == 0 {
	//			fmt.Println("PIN")
	//
	//		} else if i%5 == 0 {
	//			fmt.Println("PAN")
	//
	//		} else if i%5 == 0 && i%3 == 0 {
	//			fmt.Println("PINPAN")
	//
	//		}
	//		fmt.Println(i)
	//	}

	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PINPAN")

		} else if i%3 == 0 {
			fmt.Println("PIN")
		} else if i%5 == 0 {
			fmt.Println("PAN")
		} else {
			fmt.Println(i)
		}

	}
}

//agora foiiiiii
