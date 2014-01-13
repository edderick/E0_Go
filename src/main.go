package main

import (
	"fmt"
	"./EncryptionEngine"
)

func main(){
	/* test case 1 

	var Kc [16]byte = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var BD_ADDR [6]byte = [6]byte{0, 0, 0, 0, 0, 0}
	var CLK26 [4]byte = [4]byte{0, 0, 0, 0}

	/* test case 2 

	var Kc [16]byte = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var BD_ADDR [6]byte = [6]byte{0, 0, 0, 0, 0, 0}
	var CLK26 [4]byte = [4]byte{0, 0, 0, 3}

	/* test case 3 

	var Kc [16]byte = [16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	var BD_ADDR [6]byte = [6]byte{255, 255, 255, 255, 255, 255}
	var CLK26 [4]byte = [4]byte{255, 255, 255, 3}

	/* test case 4 */

	var Kc [16]byte = [16]byte{0x21, 0x87, 0xF0, 0x4a, 0xba, 0x90, 0x31, 0xd0, 0x78, 0x0d, 0x4c, 0x53, 0xe0, 0x15, 0x3a, 0x63}
	var BD_ADDR [6]byte = [6]byte{0x2c, 0x7f, 0x94, 0x56, 0x0f, 0x1b}
	var CLK26 [4]byte = [4]byte{0x5f, 0x1a, 0x00, 0x02}

	/* end test cases */

	var bytesReq int = 16;
	fmt.Println(EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26,
			bytesReq))
}
