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
   actual := sm.step(0, 0, 0, 0)

   if expected != actual {
        t.Error("Failed for all inputs as zero, double step.")
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
