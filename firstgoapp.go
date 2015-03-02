package main

import (
  "fmt"
  "time"
  "math"
  )
  
func add(x, y int) int {
  return x + y
  }
  
func swap(a, b string) (string, string){
  return b, a
  }

func split(sum int) (x, y int) {
  x = sum * 4/9
  y = sum - x
  return
  }
  
var c, python, java = true, false, "no!"

func main() {
  fmt.Println("What up")
  fmt.Println("The time is", time.Now())
//  fmt.Println("May favorite number is", math.rand.Intn(10))
  fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
  fmt.Println(math.Pi)
  fmt.Println(add(4,5))
  a, b := swap("Hello","World")
  fmt.Println(a,b)
  fmt.Println(split(17))
  i, j := 1, 2
  k := 3
  fmt.Println(i, j, k, c, python, java)
  sum := 0
  for i := 0; i < 15; i++ {
    sum += i
    }
  fmt.Println(sum)
  }