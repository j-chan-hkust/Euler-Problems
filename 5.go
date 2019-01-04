package main
import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?


func main() {

	maxNum = 20
	// i mean the question is just a simple case of finding max number of times each prime factor appears in any one number in the numbers 1-20
	// eg, 2 appears at max 4 times, as 16 = 2x2x2x2
	var primeAppeareanceRate map[int]int = make(map[int]int) // we initialize a map mapping prime to appeareance rate

	primes := [1]int{2} // we initialize for the first prime
	primeAppeareanceRate[2] = 1

	// im pretty sure there is a better data structure, but im not good with go

	for i := 3; i<=maxNum; i++ {
		for j := 0; i != 1; j++{

		}
		//for each number, see if its prime if it's prime, then we will add it to the map
		//if its not prime, we will see how many 
	}

    fmt.Println("hello world")
}