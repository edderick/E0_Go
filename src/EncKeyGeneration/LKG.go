package main

import (
	"fmt"
	"math/big"
)

func E(X [12]uint8, L int) [16]uint8 {
	var Y [16]uint8
	for i := 0; i < 16; i++{
		Y[i] = X[i%L]
	}
	return Y
}

func PHT(in [2]uint8) [2]uint8 {
	var out [2]uint8
	out[0] = uint8(uint(2*in[0] + in[1])%256)
	out[1] = uint8(uint(in[0] + in[1])%256)
	return out
}

func e(i uint8) uint8{
	var I *big.Int = big.NewInt(int64(i))
	I.Exp(big.NewInt(45), I, big.NewInt(257))
	I.Mod(I, big.NewInt(256))
	return uint8(I.Uint64())
}

func l(i uint8) uint8{
	var j uint8
	for j = 0; uint(j) < 256; j++{
		if (e(j) == i){
			return j
		}
	}
	return 0
}

func main(){
	var COF [12]uint8 = [12]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(COF)
	fmt.Println(E(COF, 12))

	var in [2]uint8 = [2]uint8{200, 250}
	fmt.Println(in)
	fmt.Println(PHT(in))

	for i := 0; i < 256; i++{
		num := e(uint8(i))
		var j uint8 = l(num)
		fmt.Printf("%d: %d, %d\n", i, j, num)
	}
}
