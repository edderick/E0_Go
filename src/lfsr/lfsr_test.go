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

    e0_l := NewE0_LFSRs()

    fmt.Println(e0_l.lfsr_1.Next())
    fmt.Println(e0_l.lfsr_2.Next())
    fmt.Println(e0_l.lfsr_3.Next())
    fmt.Println(e0_l.lfsr_4.Next())
}
