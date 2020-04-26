package helpers

import (
	"fmt"
	"os"
)

func DD(any interface{}) {
	fmt.Printf("El valor de la variable es %v \n tipo %T\n", any, any)
	os.Exit(2)
}

func Exit(e error) {
	fmt.Println(e)
	os.Exit(2)
}
