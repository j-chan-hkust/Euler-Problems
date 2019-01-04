package main
import "fmt"
import "time"

// The sum of the squares of the first ten natural numbers is,
//
// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

// there probably IS a O(logn) way (or better) to achieve this, but i really cant be bothered

func main(){
  start := time.Now()
  var input int = 100
  var sum1 int = 0
  for i:=0;i<input;i++{
    sum1 += i*i
  }

  var sum2 int = (input+1)*input/2
  sum2 = sum2*sum2

  fmt.Println(sum2-sum1)

  elapsed := time.Since(start)
  fmt.Println("time taken is", elapsed)
}
