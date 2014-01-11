package lfsr

type LFSR struct {
    switch_closed bool
    num_bits_shifted int 

    values []bool 
    taps []int
}

func NewLFSR(length int, taps []int) *LFSR {
    l := new(LFSR)

    l.switch_closed = false
    l.num_bits_shifted = 0
    
    l.taps = taps
    l.values = make([]bool, length)

    return l
}

func (l LFSR) Shift(val bool) bool {
    out := l.values[len(l.values) - 1]
    
    for i, v := range(l.values) {
        l.values[i] = val
        val = v 
    }
    
    l.num_bits_shifted++ 
    if !l.switch_closed && (l.num_bits_shifted >= len(l.values)) {
        l.switch_closed = true
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

func (l LFSR) NextBit(bit bool) bool {
    if l.switch_closed {
        return l.Shift(bool_xor(l.feedback(), bit))
    } else {
        return l.Shift(bit)
    }
}

func getBit(val byte, i uint) bool {
    return (val & (1 << i)) != 0
}

func (l LFSR) NextByte(val byte) byte {
    var out byte
    out = 0

    for i := uint(0); i < 8; i++ {
        bit := getBit(val, i) 
        
        out_bit := l.NextBit(bit)
        if out_bit {
            out = out | (1 << 7)
        }
        out = out >> 1
    }

    return out
}
