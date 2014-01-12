package main

import (
	"fmt"
	"./EncryptionEngine"
)

func main(){
	var Kc [16]byte = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var BD_ADDR [6]byte = [6]byte{0, 0, 0, 0, 0, 0}
	var CLK26 [4]byte = [4]byte{0, 0, 0, 0}
	var bytesReq int = 8;
	fmt.Println(EncryptionEngine.GetPayloadKey(Kc, BD_ADDR, CLK26,
			bytesReq))
}
