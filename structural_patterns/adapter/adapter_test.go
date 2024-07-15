package adapter

import "testing"

func TestAdapter_Request(t *testing.T) {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)
	if adapter.Request() != "specific request" {
		t.Error("Adapter does not return expected value")
	}
}
