package builder

import (
	"testing"
)

func TestHouseBuilder_SetWindowType(t *testing.T) {
	b := &HouseBuilder{}

	// Test case 1: Set window type to "wooden"
	b.SetWindowType("wooden")
	if b.h.windowType != "wooden" {
		t.Errorf("Expected window type to be 'wooden', got '%s'", b.h.windowType)
	}

	// Test case 2: Set window type to an empty string
	b.SetWindowType("")
	if b.h.windowType != "" {
		t.Errorf("Expected window type to be an empty string, got '%s'", b.h.windowType)
	}

	// Test case 3: Set window type to a long string
	longString := "abcdefghijklmnopqrstuvwxyz"
	b.SetWindowType(longString)
	if len(b.h.windowType) != len(longString) {
		t.Errorf("Expected window type length to be %d, got %d", len(longString), len(b.h.windowType))
	}
}

func TestSetDoorType(t *testing.T) {
	builder := &HouseBuilder{}

	// Setting a new door type
	newDoorType := "Wooden"
	builder.SetDoorType(newDoorType)
	if builder.h.doorType != newDoorType {
		t.Errorf("Expected door type to be %s, but got %s", newDoorType, builder.h.doorType)
	}

	// Setting an empty door type
	emptyDoorType := ""
	builder.SetDoorType(emptyDoorType)
	if builder.h.doorType != emptyDoorType {
		t.Errorf("Expected door type to be %s, but got %s", emptyDoorType, builder.h.doorType)
	}
}

func TestHouseBuilder_SetFloor(t *testing.T) {
	tests := []struct {
		name  string
		floor int
		want  int
	}{
		{
			name:  "Set floor to 1",
			floor: 1,
			want:  1,
		},
		{
			name:  "Set floor to 2",
			floor: 2,
			want:  2,
		},
		{
			name:  "Set floor to -1",
			floor: -1,
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HouseBuilder{}
			h.SetFloor(tt.floor)

			if h.h.floor != tt.want {
				t.Errorf("Expected floor to be %d, but got %d", tt.want, h.h.floor)
			}
		})
	}
}
