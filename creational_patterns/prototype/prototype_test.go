package prototype

import (
	"reflect"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	p := ConcretePrototype{Field1: "test", Field2: 123}
	clone := p.DeepCopy().(*ConcretePrototype)

	if !reflect.DeepEqual(p, *clone) {
		t.Errorf("DeepCopy did not correctly copy the object.")
	}

	if &p == clone {
		t.Errorf("DeepCopy did not return a new instance.")
	}

	// Test handling nil input
	var nilP *ConcretePrototype
	clone = nilP.DeepCopy().(*ConcretePrototype)
	if clone == nil {
		t.Errorf("DeepCopy did not handle nil input gracefully.")
	}
}

// func TestClone(t *testing.T) {
// 	// Test that the clone has the same values as the original
// 	original := &ConcretePrototype{
// 		Field1: "test",
// 		Field2: 42,
// 	}
// 	clone := original.Clone()
// 	if clone.Field1 != original.Field1 {
// 		t.Errorf("Clone has incorrect Field1 value: got %s, want %s", clone.Field1, original.Field1)
// 	}
// 	if clone.Field2 != original.Field2 {
// 		t.Errorf("Clone has incorrect Field2 value: got %d, want %d", clone.Field2, original.Field2)
// 	}

// 	// Test that modifying the clone does not modify the original
// 	clone.Field1 = "changed"
// 	if original.Field1 == clone.Field1 {
// 		t.Errorf("Changing clone modified original: got %s, want %s", original.Field1, "test")
// 	}

// 	// Test that the clone is a separate object
// 	if original == clone {
// 		t.Errorf("Clone is not a separate object: got %p, want %p", original, clone)
// 	}
// }
