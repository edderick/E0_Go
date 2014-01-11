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



//func init(e0_lfsrs *E0_LFSRs, Kc []byte, bd_addr []byte, clock []byte) {
//}
