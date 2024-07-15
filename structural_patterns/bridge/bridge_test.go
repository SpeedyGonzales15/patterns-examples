package bridge

import "testing"

func TestConcreteImplementorA_OperationImpl(t *testing.T) {
	c := &ConcreteImplementorA{}
	expected := "ConcreteImplementorA OperationImpl"
	result := c.OperationImpl()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestConcreteImplementorB_OperationImpl(t *testing.T) {
	c := &ConcreteImplementorB{}
	expected := "ConcreteImplementorB OperationImpl"
	result := c.OperationImpl()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
