package main

import (
	"fmt"
    "encoding/binary"
	"./EncryptionEngine"
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

func CLK26_to_clock(CLK26 [4]byte) uint32 {
    for i := 0; i < 4; i++ {
        CLK26[i] = reverse[CLK26[i]]
    }
    return binary.LittleEndian.Uint32(CLK26[:])
}

func main(){
	/* Test Case 1 */ 
	Kc := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	BD_ADDR := [6]byte{0, 0, 0, 0, 0, 0}
	CLK26 := [4]byte{0, 0, 0, 0}
    clock := CLK26_to_clock(CLK26)

    test_1_expected := [16]byte{70, 105, 78, 97, 147, 52, 92, 135, 113, 24, 148, 146, 27, 183, 141, 160}
    test_1_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, clock, 16)
	
    fmt.Println("Test 1 Expected: ", test_1_expected)
    fmt.Println("Test 1 Actual:   ", test_1_actual)
    fmt.Println() 

    /* Test Case 2 */
	Kc = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	BD_ADDR = [6]byte{0, 0, 0, 0, 0, 0}
	CLK26 = [4]byte{0, 0, 0, 3}
    clock = CLK26_to_clock(CLK26)

    test_2_expected := [16]byte{140, 195, 27, 57, 5, 214, 42, 76, 54, 141, 90, 210, 74, 67, 54, 74}
    test_2_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, clock, 16)
	
    fmt.Println("Test 2 Expected: ", test_2_expected)
    fmt.Println("Test 2 Actual:   ", test_2_actual)
    fmt.Println() 
	
    /* Test Case 3 */ 
	Kc = [16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	BD_ADDR = [6]byte{255, 255, 255, 255, 255, 255}
	CLK26 = [4]byte{255, 255, 255, 3}
    clock = CLK26_to_clock(CLK26)

    test_3_expected := [16]byte{139, 255, 253, 169, 141, 214, 249, 121, 194, 188, 101, 88, 83, 23, 132, 78}
    test_3_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, clock, 16)
	
    fmt.Println("Test 3 Expected: ", test_3_expected)
    fmt.Println("Test 3 Actual:   ", test_3_actual)
    fmt.Println() 

	/* Test Case 4 */
	Kc = [16]byte{0x21, 0x87, 0xF0, 0x4a, 0xba, 0x90, 0x31, 0xd0, 0x78, 0x0d, 0x4c, 0x53, 0xe0, 0x15, 0x3a, 0x63}
	BD_ADDR = [6]byte{0x2c, 0x7f, 0x94, 0x56, 0x0f, 0x1b}
	CLK26 = [4]byte{0x5f, 0x1a, 0x00, 0x02}
    clock = CLK26_to_clock(CLK26)

    test_4_expected := [16]byte{41, 153, 246, 7, 253, 224, 46, 164, 204, 156, 27, 133, 3, 165, 148, 41}
    test_4_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, clock, 16)
	
    fmt.Println("Test 4 Expected: ", test_4_expected)
    fmt.Println("Test 4 Actual:   ", test_4_actual)
    fmt.Println() 
}
