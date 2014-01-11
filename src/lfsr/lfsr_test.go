package lfsr 

import "testing"
import "fmt"

func Test_Simple(t *testing.T) {
    taps := []int{1, 2, 3, 4, 5}

    l := NewLFSR(10, taps)
  
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

func Test_Four(t *testing.T) {
    tap_1 := []int{25, 20, 12, 8}
    tap_2 := []int{31, 24, 16, 12}
    tap_3 := []int{33, 28, 24, 4}
    tap_4 := []int{39, 36, 20, 4}

    lfsr_1 := NewLFSR(25, tap_1)
    lfsr_2 := NewLFSR(31, tap_2)
    lfsr_3 := NewLFSR(33, tap_3)
    lfsr_4 := NewLFSR(39, tap_4)

    fmt.Println(lfsr_1.Next())
    fmt.Println(lfsr_2.Next())
    fmt.Println(lfsr_3.Next())
    fmt.Println(lfsr_4.Next())
}
