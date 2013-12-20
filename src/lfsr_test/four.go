package main 

import "fmt"
import "lfsr"

func main() {
    tap_1 := []int{25, 20, 12, 8}
    tap_2 := []int(31, 24, 16, 12)
    tap_3 := []int(33, 28, 24, 4)
    tap_4 := []int(39, 36, 20, 4)

    lfsr_1 := lfsr.NewLFSR(25, tap_1)
    lfsr_2 := lfsr.NewLFSR(31, tap_2)
    lfsr_3 := lfsr.NewLFSR(33, tap_3)
    lfsr_4 := lfsr.NewLFSR(39, tap_4)
    

}
