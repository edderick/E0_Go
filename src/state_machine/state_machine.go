package state_machine 

type StateMachine struct {

    c_t_minus_one int
    c_t int

}

func T1(in int) int {
    return in
}

//XXX: It would probably be more readable just to do this as a case...
func T2(in int) int {
    in_x0 := in & 1
    in_x1 := (in >> 1) & 1

    out_x1 := in_x0
    out_x0 := in_x0 ^ in_x1

    out := out_x0 | (out_x1 << 1)
    return out
}

func (sm *StateMachine) step(x1, x2, x3, x4 int) bool {
    y := x1 + x2 + x3 + x4

    s := (y + sm.c_t) / 2

    c_t_plus_one := T1(sm.c_t) ^ T2(sm.c_t_minus_one) ^ s

    z := (x1 ^ x2 ^ x3 ^x4 ^ (sm.c_t & 1)) == 1 

    sm.c_t_minus_one = sm.c_t 
    sm.c_t = c_t_plus_one 

    return z
}

