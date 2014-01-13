package state_machine

type StateMachine struct {

    C_t_minus_one int
    C_t int
    C_t_plus_one int

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

func (sm *StateMachine) FireEDC(){

    sm.C_t_minus_one = sm.C_t
    sm.C_t = sm.C_t_plus_one
}

func (sm *StateMachine) Step(x [4]int) bool {

    y := x[0] + x[1] + x[2] + x[3]

    s := (y + sm.C_t) / 2

    sm.C_t_plus_one = T1(sm.C_t) ^ T2(sm.C_t_minus_one) ^ s

    z := (x[0] ^ x[1] ^ x[2] ^x[3] ^ (sm.C_t & 1)) == 1

    return z
}

func (sm *StateMachine) Reset(){
	sm.C_t = 0
	sm.C_t_minus_one = 0
}
