package EncryptionEngine

import (
	"fmt"
	"./lfsr"
	"./state_machine"
)

func GetKeyStream(Kc [16]byte, BD_ADDR [6]byte, CLK26 [4]byte,
		bytesReq int) []byte{
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
			uint64((CLK26[3] & 1))
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
			uint64((CLK26[3] & 2) >> 1)
	inputs[3] = (uint64(BD_ADDR[5]) << 47) |
			(uint64(BD_ADDR[1]) << 39) |
			(uint64(Kc[15]) << 31) |
			(uint64(Kc[11]) << 23) |
			(uint64(Kc[7]) << 15) |
			(uint64(Kc[3]) << 7) |
			(uint64(CLK26[0] & 240) >> 1) | 7

	var sm state_machine.StateMachine = state_machine.StateMachine{0, 0, 0}
	var lfsrout [4]int
	var z bool
	for t := 0; t < 39; t++{
		ClockLFSRs(lfsrs, &inputs)
		lfsrout = GetLFSROut(lfsrs)
		sm.FireEDC()
		sm.Reset()
		z = sm.Step(lfsrout)
	}

	var Z [16]byte
	for t := 39; t < 239; t++{
		ClockLFSRs(lfsrs, &inputs)
		lfsrout = GetLFSROut(lfsrs)
		sm.FireEDC()
		z = sm.Step(lfsrout)
		if z && t >= 111{
			Z[(t-111)/8] = Z[(t-111)/8] | (1 << ((uint(t) - 111) % 8))
		}
	}
	var pinputs [4]uint64
	pinputs[0] = uint64(Z[0]) |
			(uint64(Z[4]) << 8) |
			(uint64(Z[8]) << 16) |
			(uint64(Z[12] & 1) << 24)
	pinputs[1] = (uint64(Z[1])) |
			(uint64(Z[5]) << 8) |
			(uint64(Z[9]) << 16) |
			(uint64(Z[12] & 254) << 23)
	pinputs[2] = (uint64(Z[2])) |
			(uint64(Z[6]) << 8) |
			(uint64(Z[10]) << 16) |
			(uint64(Z[13]) << 24) |
			(uint64(Z[15] & 1) << 32)
	pinputs[3] = (uint64(Z[3])) |
			(uint64(Z[7]) << 8) |
			(uint64(Z[11]) << 16) |
			(uint64(Z[14]) << 24) |
			(uint64(Z[15] & 254) << 31)

	for x := 0; x < 4; x++{
		lfsrs[x].ParallelLoad(pinputs[x])
	}

	var KeyStream = make([]byte, bytesReq)
	lfsrout = GetLFSROut(lfsrs)
	z = sm.Step(lfsrout)
	if z{
        KeyStream[0] = 1 << 7
	} 
    
    for t := 240; t < 239 + bytesReq * 8; t++{
		ClockLFSRs(lfsrs, &inputs)
		lfsrout = GetLFSROut(lfsrs)
		sm.FireEDC()
		z = sm.Step(lfsrout)
		if z{
			KeyStream[(t-239)/8] = KeyStream[(t-239)/8] | (1 << (7 -(uint(t+1) % 8)))
		}
	}

	return KeyStream
}

func PrintState(t int, lfsrs [4]*lfsr.LFSR, sm *state_machine.StateMachine){
	fmt.Printf("%d ", t)
	for x := 0; x < 4; x++{
		fmt.Printf("% X ", lfsrs[x].GetContents())
	}

	fmt.Printf("%.2b %.2b %.2b\n", sm.C_t_plus_one, sm.C_t, sm.C_t_minus_one)
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

func Encrypt(pt, keyStream []byte) []byte{
    ct := make([]byte, len(pt))

    for i := 0; i < len(pt); i++ {
        ct[i] = pt[i] ^ keyStream[i]
    }

    return ct 
}
