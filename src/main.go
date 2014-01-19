package main

import (
	"fmt"
	"./EncryptionEngine"
)

func main(){
	/* Test Case 1 */ 
	Kc := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	BD_ADDR := [6]byte{0, 0, 0, 0, 0, 0}
	CLK26 := [4]byte{0, 0, 0, 0}

    test_1_expected := [16]byte{70, 105, 78, 97, 147, 52, 92, 135, 113, 24, 148, 146, 27, 183, 141, 160}
    test_1_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 1 Expected: ", test_1_expected)
    fmt.Println("Test 1 Actual:   ", test_1_actual)
    fmt.Println() 

    /* Test Case 2 */
	Kc = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	BD_ADDR = [6]byte{0, 0, 0, 0, 0, 0}
	CLK26 = [4]byte{0, 0, 0, 3}

    test_2_expected := [16]byte{140, 195, 27, 57, 5, 214, 42, 76, 54, 141, 90, 210, 74, 67, 54, 74}
    test_2_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 2 Expected: ", test_2_expected)
    fmt.Println("Test 2 Actual:   ", test_2_actual)
    fmt.Println() 
	
    /* Test Case 3 */ 
	Kc = [16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	BD_ADDR = [6]byte{255, 255, 255, 255, 255, 255}
	CLK26 = [4]byte{255, 255, 255, 3}

    test_3_expected := [16]byte{139, 255, 253, 169, 141, 214, 249, 121, 194, 188, 101, 88, 83, 23, 132, 78}
    test_3_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 3 Expected: ", test_3_expected)
    fmt.Println("Test 3 Actual:   ", test_3_actual)
    fmt.Println() 

	/* Test Case 4 */
	Kc = [16]byte{0x21, 0x87, 0xF0, 0x4a, 0xba, 0x90, 0x31, 0xd0, 0x78, 0x0d, 0x4c, 0x53, 0xe0, 0x15, 0x3a, 0x63}
	BD_ADDR = [6]byte{0x2c, 0x7f, 0x94, 0x56, 0x0f, 0x1b}
	CLK26 = [4]byte{0x5f, 0x1a, 0x00, 0x02}

    test_4_expected := [16]byte{41, 153, 246, 7, 253, 224, 46, 164, 204, 156, 27, 133, 3, 165, 148, 41}
    test_4_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 4 Expected: ", test_4_expected)
    fmt.Println("Test 4 Actual:   ", test_4_actual)
    fmt.Println() 
}
