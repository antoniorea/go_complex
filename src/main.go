/*    
    Este programa suma y resta dos numeros complejos en coord. polares
    los argumentos deben estar expresados en grados
$ ./complex01_go 12.041 48.36 8.60 54.46 

z1 + z2 :       El modulo de z es :20.61 el argumento z en grados es :50.90     (13.00,16.00)
z1 - z2 :       El modulo de z es :3.61 el argumento z en grados es :33.69      (3.00,2.00)
*/
package main

import (
	"os"
	"strconv"
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	
	var numbers []float64
	for _, a := range os.Args[1:]{
		i, err := strconv.ParseFloat(a, 64)
		numbers = append(numbers, i)
		if err != nil {
			panic(fmt.Sprintf("Invalid Value : %v", err))
		}
			
	}

	z1 := cmplx.Rect(numbers[0], (numbers[1]  * math.Pi / 180.0))
	z2 := cmplx.Rect(numbers[2], (numbers[3] * math.Pi / 180.0))

	suma := z1 + z2
	resta := z1 -z2

	r1, theta1 := cmplx.Polar(suma)	
	fmt.Printf("z1 + z2 :    El modulo de z es :%.2f el argumento z en grados es :%.2f\n",r1, theta1 * 180.0/math.Pi )
	fmt.Printf("suma: %.2f, + %.2f j\n", real(suma), imag(suma))

	r2, theta2 := cmplx.Polar(resta)	
	fmt.Printf("z1 - z2 :    El modulo de z es :%.2f el argumento z en grados es :%.2f\n",r2, theta2 * 180.0/math.Pi )
	fmt.Printf("resta: %.2f, + %.2f j\n", real(resta), imag(resta))


}

