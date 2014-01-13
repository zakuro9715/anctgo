package anctgo

import (
  "testing"
)

var initialTests = []struct {
  in  Department
  out string
}{
  {Mechanical, "M"},
  {Electrical, "E"},
  {Civil, "C"},
  {Architecture, "A"},
  {MechanicalAndElectronic, "ME"},
  {ArchitectureAndCivil, "AC"},
  {M, "M"},
  {E, "E"},
  {C, "C"},
  {A, "A"},
  {ME, "ME"},
  {AC, "AC"},
}

func TestInitial(t *testing.T) {
  for _, it := range initialTests {
    if out := it.in.Initial(); out != it.out {
      t.Errorf("(%s).Initial->(%s), (expected %s)", it.in, it.out, out)
    }
  }
}
