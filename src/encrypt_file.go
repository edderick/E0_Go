package main 

import (
    "io/ioutil"
    "./EncryptionEngine"
    "flag"
)

func main() {

    in_file := flag.String("in", "test.txt", "Name of the input file")
    out_file := flag.String("out", "cipher.txt", "Name of the output file")

    flag.Parse()

    /* Encypt File */
    pt, _ := ioutil.ReadFile(*in_file)

    Kc := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    BD_ADDR := [6]byte{0, 0, 0, 0, 0, 0}
    CLK26 := [4]byte{0, 0, 0, 0}

    keyStream := EncryptionEngine.GetKeyStream(Kc, BD_ADDR, CLK26, len(pt))

    ct := make([]byte, len(pt))

    for i := 0; i < len(pt); i++ {
        ct[i] = pt[i] ^ keyStream[i]
    }

    ioutil.WriteFile(*out_file, ct, 0666)
}
