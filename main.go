package main

import (
	"fmt"
	"github.com/nelsongp/multiCoder/factory"
	"os"
)

func main(){
	var n string
	fmt.Println("Digite un nombre: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Printf("error al leer el nombre: %v", err)
		os.Exit(1)
	}

	resp := factory.Factory(n)

	fmt.Println(resp)
}
