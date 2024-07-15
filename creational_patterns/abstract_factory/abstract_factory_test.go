package abstractfactory

import "testing"

func TestLinuxFactory_CreateButton(t *testing.T) {
	factory := &LinuxFactory{}

	button := factory.CreateButton()

	if _, ok := button.(*LinuxButton); !ok {
		t.Errorf("Expected WindowsButton, got %T", button)
	}
}

func TestMacOSFactory_CreateButton(t *testing.T) {
	factory := &MacOSFactory{}

	button := factory.CreateButton()

	if _, ok := button.(*MacOSButton); !ok {
		t.Errorf("Expected WindowsButton, got %T", button)
	}
}

func TestWindowsFactory_CreateButton(t *testing.T) {
	factory := &WindowsFactory{}

	button := factory.CreateButton()

	if _, ok := button.(*WindowsButton); !ok {
		t.Errorf("Expected WindowsButton, got %T", button)
	}
}
