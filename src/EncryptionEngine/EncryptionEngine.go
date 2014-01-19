package EncryptionEngine

import (
	"fmt"
    "encoding/binary"
	"./lfsr"
	"./state_machine"
)

var reverse = [256]byte{
    0, 128, 64, 192, 32, 160, 96, 224,
    16, 144, 80, 208, 48, 176, 112, 240,
    8, 136, 72, 200, 40, 168, 104, 232,
    24, 152, 88, 216, 56, 184, 120, 248,
    4, 132, 68, 196, 36, 164, 100, 228,
    20, 148, 84, 212, 52, 180, 116, 244,
    12, 140, 76, 204, 44, 172, 108, 236,
    28, 156, 92, 220, 60, 188, 124, 252,
    2, 130, 66, 194, 34, 162, 98, 226,
    18, 146, 82, 210, 50, 178, 114, 242,
    10, 138, 74, 202, 42, 170, 106, 234,
    26, 154, 90, 218, 58, 186, 122, 250,
    6, 134, 70, 198, 38, 166, 102, 230,
    22, 150, 86, 214, 54, 182, 118, 246,
    14, 142, 78, 206, 46, 174, 110, 238,
    30, 158, 94, 222, 62, 190, 126, 254,
    1, 129, 65, 193, 33, 161, 97, 225,
    17, 145, 81, 209, 49, 177, 113, 241,
    9, 137, 73, 201, 41, 169, 105, 233,
    25, 153, 89, 217, 57, 185, 121, 249,
    5, 133, 69, 197, 37, 165, 101, 229,
    21, 149, 85, 213, 53, 181, 117, 245,
    13, 141, 77, 205, 45, 173, 109, 237,
    29, 157, 93, 221, 61, 189, 125, 253,
    3, 131, 67, 195, 35, 163, 99, 227,
    19, 147, 83, 211, 51, 179, 115, 243,
    11, 139, 75, 203, 43, 171, 107, 235,
    27, 155, 91, 219, 59, 187, 123, 251,
    7, 135, 71, 199, 39, 167, 103, 231,
    23, 151, 87, 215, 55, 183, 119, 247,
    15, 143, 79, 207, 47, 175, 111, 239,
    31, 159, 95, 223, 63, 191, 127, 255,
}

func GetKeyStream(Kc [16]byte, BD_ADDR [6]byte, clock uint32,
		bytesReq int) []byte{
	var lfsrs [4]*lfsr.LFSR
	lfsrs[0] = lfsr.NewLFSR(25, []int{8, 12, 20, 25}, 24)
	lfsrs[1] = lfsr.NewLFSR(31, []int{12, 16, 24, 31}, 24)
	lfsrs[2] = lfsr.NewLFSR(33, []int{4, 24, 28, 33}, 32)
	lfsrs[3] = lfsr.NewLFSR(39, []int{4, 28, 36, 39}, 32)

    var CLK26 [4]byte
    binary.LittleEndian.PutUint32(CLK26[:], clock)

    for i := 0; i < 4; i++ {
        CLK26[i] = reverse[CLK26[i]]
    }  

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
