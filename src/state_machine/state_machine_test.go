package state_machine

import "testing"

func Test_Zero(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := false
   actual := sm.step(0, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for all inputs as zero, single step.")
   }
}

func Test_DoubleZero(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := false
   sm.step(0, 0, 0, 0)
   actual := sm.step(0, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for all inputs as zero, double step.")
   }
}

func Test_MultipleZero(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := false
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   sm.step(0, 0, 0, 0)
   actual := sm.step(0, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for all inputs as zero, many step.")
   }
}

func Test_One(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := true
   actual := sm.step(1, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, single step.")
   }
   
   expected = true
   actual = sm.step(0, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, single step.")
   }
   
   expected = true
   actual = sm.step(0, 0, 1, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, single step.")
   }
   
   expected = true
   actual = sm.step(0, 0, 0, 1)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, single step.")
   }
}

// When we only have a single bit set, the value of c will never change 
// This is thanks to the /2 component 
func Test_DoubleOne(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := true
   sm.step(1, 0, 0, 0)
   actual := sm.step(1, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, double step.")
   }
   
   expected = true
   sm.step(0, 1, 0, 0)
   actual = sm.step(0, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, double step.")
   }
   
   expected = true
   sm.step(0, 0, 1, 0)
   actual = sm.step(0, 0, 1, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, double step.")
   }
   
   expected = true
   sm.step(0, 0, 0, 1)
   actual = sm.step(0, 0, 0, 1)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, double step.")
   }
}

func Test_MultipleOne(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := true
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   actual := sm.step(1, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, many step.")
   }
   
   expected = true
   sm.step(0, 1, 0, 0)
   sm.step(0, 1, 0, 0)
   sm.step(0, 1, 0, 0)
   sm.step(0, 1, 0, 0)
   sm.step(0, 1, 0, 0)
   actual = sm.step(0, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, many step.")
   }
   
   expected = true
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 1, 0)
   actual = sm.step(0, 0, 1, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, many step.")
   }
   
   expected = true
   sm.step(0, 0, 0, 1)
   sm.step(0, 0, 0, 1)
   sm.step(0, 0, 0, 1)
   sm.step(0, 0, 0, 1)
   sm.step(0, 0, 0, 1)
   actual = sm.step(0, 0, 0, 1)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, many step.")
   }
}

// The order of the bits doesn't actually matter...
func Test_PermutingOne(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := true
   sm.step(1, 0, 0, 0)
   sm.step(0, 1, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 0, 1)
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   sm.step(1, 0, 0, 0)
   sm.step(0, 1, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 0, 1)
   sm.step(0, 1, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 0, 1)
   sm.step(0, 1, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 0, 1)
   sm.step(0, 1, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 0, 1)
   sm.step(0, 1, 0, 0)
   sm.step(0, 0, 1, 0)
   sm.step(0, 0, 0, 1)
   actual := sm.step(1, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for a single input as one, the rest zero, many step, random permutations.")
   }
}

func Test_Two(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := false
   actual := sm.step(1, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a two inputs as one, the rest zero, single step.")
   }
}

func Test_DoubleTwo(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := true
   sm.step(1, 1, 0, 0)
   actual := sm.step(1, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a two inputs as one, the rest zero, double step.")
   }
}

func Test_MultipleTwo(t *testing.T) {
   sm := StateMachine{0, 0}

   expected := false
   sm.step(1, 1, 0, 0)
   sm.step(1, 1, 0, 0)
   actual := sm.step(1, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a two inputs as one, the rest zero, three step.")
   }
   
   expected = false
   actual = sm.step(1, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a two inputs as one, the rest zero, four step.")
   }
   
   expected = false
   actual = sm.step(1, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a two inputs as one, the rest zero, five step.")
   }

   expected = false
   actual = sm.step(1, 1, 0, 0)

   if expected != actual {
        t.Error("Failed for a two inputs as one, the rest zero, six step.")
   }
}
