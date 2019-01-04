package main
import "fmt"
import "time"


func main(){
  start := time.Now()

  elapsed := time.Since(start)
  fmt.Println("time taken is", elapsed)
}
