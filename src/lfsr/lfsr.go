package lfsr

type LFSR struct {
    values []bool 
    taps []int
}

func NewLFSR(length int, taps []int) *LFSR {
    l := new(LFSR)

    l.taps = taps
    l.values = make([]bool, length)

    return l
}

func (l LFSR) Shift(val bool) bool{
    out := l.values[len(l.values) - 1]
    
    for i, v := range(l.values) {
        l.values[i] = val
        val = v 
    }
    
    return out
}

//I have no idea why there isn't already a boolean XOR...
func bool_xor(A, B bool) bool {
    return (A && !B) || (!A && B)
}

func (l LFSR) feedback() bool {
    val := false

    for _, v := range(l.taps) {
        val = bool_xor(val, l.values[v - 1])
    }

    return val
}

func (l LFSR) Next() bool {
    return l.Shift(l.feedback())
}


