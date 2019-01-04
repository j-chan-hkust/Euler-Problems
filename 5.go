package main
import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

// time complexity: first for loop iterates n times, from a quora post, we have prime density approaches 0
// thus I assume the nested loop can be seen as having a fixed time complexity, O(1)
// the further nested loop has log(n) complexity
// total complexity of first loop is O(nlogn)
// second loop is O(1) complexity, by my assumption

func main() {

	var maxNum int = 20
	// i mean the question is just a simple case of finding max number of times each prime factor appears in any one number in the numbers 1-20
	// eg, 2 appears at max 4 times, as 16 = 2x2x2x2
	var primeAppeareanceRate map[int]int = make(map[int]int) // we initialize a map mapping prime to appeareance rate

	primes := []int{2} // we initialize for the first prime
	primeAppeareanceRate[2] = 1

	// im pretty sure there is a better data structure, but im not good with go

	//for each number, see if its prime if it's prime, then we will add it to the map
	//if its not prime, we will see how many times each prime has occured and compare with previous

	for i := 3; i<=maxNum; i++ {
		var temp int = i
		for _, prime := range primes{
			var count int = 0
			for temp%prime == 0 {
				//it is divisible by a smaller prime, thus is not prime
				count++
				temp /= prime
				//continuously divide until not divisible
			}
			if primeAppeareanceRate[prime] < count{
				primeAppeareanceRate[prime] = count
			}
		}

		if temp != 1 { //if it were divisible by smaller primes, we would be left with only 1 by now
			primeAppeareanceRate[temp] = 1
			primes = append(primes, temp)
			//we have added new prime to the list
		}
	}

	//now we times all the primes
	var output int = 1
	for _, prime:= range primes{
		for i:=0; i<primeAppeareanceRate[prime]; i++{
			output *= prime
		}
	}

    fmt.Println(output)
}
