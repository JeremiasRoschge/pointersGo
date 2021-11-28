package main

import (
	"fmt"
)

func main()  {
	/*
		¿Qué es un puntero?

		1. Es una dirección de memoria
		2. En lugar del valor, tenemos la dirección en la que esta el valor
		3. X,Y => asd123d => 5
		4. X => 6
		5 Y ¿? => 6 
		*T => *int *string *Struct
		Valor zero => nil
	*/
	var x,y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)


}