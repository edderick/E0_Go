package main

import (
	"fmt"
	"./lfsr"
	"./state_machine"
)

func main(){
	var Kc [16]byte = [16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	var BD_ADDR [6]byte = [6]byte{255, 255, 255, 255, 255, 255}
	var CLK26 [4]byte = [4]byte{255, 255, 255, 3}

	var lfsrs [4]*lfsr.LFSR
	lfsrs[0] = lfsr.NewLFSR(25, []int{8, 12, 20, 25}, 24)
	lfsrs[1] = lfsr.NewLFSR(31, []int{12, 16, 24, 31}, 24)
	lfsrs[2] = lfsr.NewLFSR(33, []int{4, 24, 28, 33}, 32)
	lfsrs[3] = lfsr.NewLFSR(39, []int{4, 28, 36, 39}, 32)

	var inputs [4]uint64
	inputs[0] = (uint64(BD_ADDR[2]) << 41) |
			(uint64(CLK26[1]) << 33) |
			(uint64(Kc[12]) << 25) |
			(uint64(Kc[8]) << 17) |
			(uint64(Kc[4]) << 9) |
			(uint64(Kc[0]) << 1) |
			uint64((CLK26[3] & 2) >> 1)
	inputs[1] = (uint64(BD_ADDR[3]) << 47) |
			(uint64(BD_ADDR[0]) << 39) |
			(uint64(Kc[13]) << 31) |
			(uint64(Kc[9]) << 23) |
			(uint64(Kc[5]) << 15) |
			(uint64(Kc[1]) << 7) |
			(uint64(CLK26[0] & 15) << 3) | 1
	inputs[2] = (uint64(BD_ADDR[4]) << 41) |
			(uint64(CLK26[2]) << 33) |
			(uint64(Kc[14]) << 25) |
			(uint64(Kc[10]) << 17) |
			(uint64(Kc[6]) << 9) |
			(uint64(Kc[2]) << 1) |
			uint64(CLK26[3] & 1)
	inputs[3] = (uint64(BD_ADDR[5]) << 47) |
			(uint64(BD_ADDR[1]) << 39) |
			(uint64(Kc[15]) << 31) |
			(uint64(Kc[11]) << 23) |
			(uint64(Kc[7]) << 15) |
			(uint64(Kc[3]) << 7) |
			(uint64(CLK26[0] & 240) >> 1) | 7

	for t := 0; t < 39; t++{
		ClockLFSRs(lfsrs, &inputs)
	}

	var sm state_machine.StateMachine = state_machine.StateMachine{0, 0}
	fmt.Println(sm)
	var Z [25]byte
	var lfsrout [4]int
	var z bool
	fmt.Println(Z, z)
	fmt.Println(lfsrout)
	for t := 39; t < 239; t++{
		ClockLFSRs(lfsrs, &inputs)
		lfsrout = GetLFSROut(lfsrs)
		z := sm.Step(lfsrout)
		fmt.Println(z) //TODO put z into Z then use as initialisation value
	}
}

func GetLFSROut(lfsrs [4]*lfsr.LFSR) [4]int {
	a := [4]int{}
	for x := 0; x < 4; x++ {
		if (lfsrs[x].GetOutput()){
			a[x] = 1
		} else {
			a[x] = 0
		}
	}
	return a
}

func ClockLFSRs(lfsrs [4]*lfsr.LFSR, inputs *[4]uint64){
	for x := 0; x < 4; x++{
		lfsrs[x].NextBit((inputs[x] & 1) == 1)
		inputs[x] = inputs[x] >> 1
	}
}
