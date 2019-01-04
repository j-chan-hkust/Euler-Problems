package main
import "fmt"
import "time"

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

//uh, i dont think there is a way other than brute forcing it lol
//time complexity is crazy I dont think theres a way to quantify it

//key takeaways: when checking for a simple condition, remember to break to save time
//looking into online lit, you can use prime sieve method, but it is memory intensive, and most importantly
// requires knowledge of approximate size of the nth prime. Don't like that

func main(){
  start := time.Now()
  var n int = 10001
  var k int = 1
  var currNum int = 3

  primes := []int{2} //initialize the first prime

  for k<n {
    isKPrime := true
    for _, prime := range primes{
      if currNum%prime == 0 {
        isKPrime = false
        break // this single line sped it up 5 times (still takes 0.5s tho)
      }
    }

    if isKPrime {
      primes = append(primes, currNum)
      k++
    }

    currNum += 2 //a very simple optimization, primes must be odd
  }

  fmt.Println(primes[len(primes)-1])

  elapsed := time.Since(start)
  fmt.Println("time taken is", elapsed)
}
