/*package main

import "fmt"

func main() {
	contador := 10

	for contador > 0 {
		fmt.Println("Contagem regressiva:", contador)
		contador--
	}

	fmt.Println("Arassou")
}

*/

package main

import "fmt"

func main() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
