package lfsr

type E0_LFSRs struct {
    lfsr_1 *LFSR
    lfsr_2 *LFSR
    lfsr_3 *LFSR
    lfsr_4 *LFSR
}

func NewE0_LFSRs() *E0_LFSRs {
    tap_1 := []int{25, 20, 12, 8}
    tap_2 := []int{31, 24, 16, 12}
    tap_3 := []int{33, 28, 24, 4}
    tap_4 := []int{39, 36, 20, 4}

    e0_lfsrs := new(E0_LFSRs)
    
    e0_lfsrs.lfsr_1 = NewLFSR(25, tap_1)
    e0_lfsrs.lfsr_2 = NewLFSR(31, tap_2)
    e0_lfsrs.lfsr_3 = NewLFSR(33, tap_3)
    e0_lfsrs.lfsr_4 = NewLFSR(39, tap_4)

    return e0_lfsrs
}

func init_e0_lfsrs(e0_lfsrs *E0_LFSRs, Kc []byte, bd_addr []byte, clk []byte) {
    init_lfsr_1(e0_lfsrs.lfsr_1, Kc, bd_addr, clk)
    init_lfsr_1(e0_lfsrs.lfsr_2, Kc, bd_addr, clk)
    init_lfsr_1(e0_lfsrs.lfsr_3, Kc, bd_addr, clk)
    init_lfsr_1(e0_lfsrs.lfsr_4, Kc, bd_addr, clk)
}

func init_lfsr_1(lfsr *LFSR, Kc []byte, bd_addr []byte, clk []byte) {
    lfsr.NextBit((clk[3] & 1) != 0) //CL24 
    lfsr.NextByte(Kc[0])
    lfsr.NextByte(Kc[4])
    lfsr.NextByte(Kc[8])
    lfsr.NextByte(Kc[12])
    lfsr.NextByte(clk[1])
    lfsr.NextByte(bd_addr[2])
}

func init_lfsr_2(lfsr *LFSR, Kc []byte, bd_addr []byte, clk []byte) {
    lfsr.NextBit(true) 
    lfsr.NextBit(false) 
    lfsr.NextBit(false) 
    lfsr.NextBit((clk[0] & (1 << 0)) != 0) //CL0L0
    lfsr.NextBit((clk[0] & (1 << 1)) != 0) //CL0L1
    lfsr.NextBit((clk[0] & (1 << 2)) != 0) //CL0L2
    lfsr.NextBit((clk[0] & (1 << 3)) != 0) //CL0L3
    lfsr.NextByte(Kc[1])
    lfsr.NextByte(Kc[5])
    lfsr.NextByte(Kc[9])
    lfsr.NextByte(Kc[13])
    lfsr.NextByte(bd_addr[0])
    lfsr.NextByte(bd_addr[3])
}

func init_lfsr_3(lfsr *LFSR, Kc []byte, bd_addr []byte, clk []byte) {
    lfsr.NextBit((clk[3] & (1 << 1)) != 0) //CL25
    lfsr.NextByte(Kc[2])
    lfsr.NextByte(Kc[6])
    lfsr.NextByte(Kc[10])
    lfsr.NextByte(Kc[14])
    lfsr.NextByte(clk[2])
    lfsr.NextByte(bd_addr[4])
}

func init_lfsr_4(lfsr *LFSR, Kc []byte, bd_addr []byte, clk []byte) {
    lfsr.NextBit(true) 
    lfsr.NextBit(true) 
    lfsr.NextBit(true) 
    lfsr.NextBit((clk[0] & (1 << 4)) != 0) //CL0U0
    lfsr.NextBit((clk[0] & (1 << 5)) != 0) //CL0U1
    lfsr.NextBit((clk[0] & (1 << 6)) != 0) //CL0U2
    lfsr.NextBit((clk[0] & (1 << 7)) != 0) //CL0U3
    lfsr.NextByte(Kc[3])
    lfsr.NextByte(Kc[7])
    lfsr.NextByte(Kc[11])
    lfsr.NextByte(Kc[15])
    lfsr.NextByte(bd_addr[1])
    lfsr.NextByte(bd_addr[5])
}
