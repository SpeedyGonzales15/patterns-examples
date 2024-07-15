package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestGetInstance_InstanceIsNil"},
		{"TestGetInstance_InstanceIsCreatedOnce"},
		{"TestGetInstance_SameInstanceReturned"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "TestGetInstance_InstanceIsNil":
				if instance != nil {
					t.Errorf("Expected instance to be nil initially")
				}
			case "TestGetInstance_InstanceIsCreatedOnce":
				GetInstance()
				GetInstance()
				if oncecall := oncecallCount(); oncecall != 1 {
					t.Errorf("Expected instance creation to be called once, got %d times", oncecall)
				}
			case "TestGetInstance_SameInstanceReturned":
				instance1 := GetInstance()
				instance2 := GetInstance()
				if instance1 != instance2 {
					t.Errorf("Expected instance returned to be the same, got different instances")
				}
			}
		})
	}
}

func oncecallCount() int {
	oncecall := 0
	once.Do(func() {
		oncecall++
	})
	return oncecall
}
