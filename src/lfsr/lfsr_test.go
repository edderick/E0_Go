package main 

import "fmt"
import "lfsr"

func main() {
    taps := []int{1, 2, 3, 4, 5}

    l := lfsr.NewLFSR(10, taps)
  
    fmt.Println(l)
    l.Shift(true)
    l.Shift(true)
    l.Shift(true)
    l.Shift(false)
    l.Shift(false)
    l.Shift(true)
    l.Shift(true)
    fmt.Println(l)

    fmt.Println(l.Next())
    fmt.Println(l.Next())
    fmt.Println(l.Next())
    fmt.Println(l.Next())
    fmt.Println(l.Next())
    fmt.Println(l.Next())

}
