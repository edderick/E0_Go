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

    test_1_expected := [16]byte{98, 150, 114, 134, 201, 44, 58, 225, 142, 24, 41, 73, 216, 237, 177, 5}
    test_1_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 1 Expected: ", test_1_expected)
    fmt.Println("Test 1 Actual:   ", test_1_actual)
    fmt.Println() 

    /* Test Case 2 */
	Kc = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	BD_ADDR = [6]byte{0, 0, 0, 0, 0, 0}
	CLK26 = [4]byte{0, 0, 0, 3}

    test_2_expected := [16]byte{49, 195, 216, 156, 160, 107, 84, 50, 108, 177, 90, 75, 82, 194, 108, 82}
    test_2_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 2 Expected: ", test_2_expected)
    fmt.Println("Test 2 Actual:   ", test_2_actual)
    fmt.Println() 
	
    /* Test Case 3 */ 
	Kc = [16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	BD_ADDR = [6]byte{255, 255, 255, 255, 255, 255}
	CLK26 = [4]byte{255, 255, 255, 3}

    test_3_expected := [16]byte{209, 255, 191, 149, 177, 107, 159, 158, 67, 61, 166, 26, 202, 232, 33, 114}
    test_3_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 3 Expected: ", test_3_expected)
    fmt.Println("Test 3 Actual:   ", test_3_actual)
    fmt.Println() 

	/* Test Case 4 */
	Kc = [16]byte{0x21, 0x87, 0xF0, 0x4a, 0xba, 0x90, 0x31, 0xd0, 0x78, 0x0d, 0x4c, 0x53, 0xe0, 0x15, 0x3a, 0x63}
	BD_ADDR = [6]byte{0x2c, 0x7f, 0x94, 0x56, 0x0f, 0x1b}
	CLK26 = [4]byte{0x5f, 0x1a, 0x00, 0x02}

    test_4_expected := [16]byte{148, 153, 111, 224, 191, 7, 116, 37, 51, 57, 216, 161, 192, 165, 41, 148}
    test_4_actual := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, 16)
	
    fmt.Println("Test 4 Expected: ", test_4_expected)
    fmt.Println("Test 4 Actual:   ", test_4_actual)
    fmt.Println() 
}
