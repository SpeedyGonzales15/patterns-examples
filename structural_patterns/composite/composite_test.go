package composite

import "testing"

func TestComposite_Operation(t *testing.T) {
	// Test case: Composite with no children
	composite := &Composite{}
	expected := "Composite with children: "
	if result := composite.Operation(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case: Composite with one child
	composite = &Composite{
		children: []Component{
			&Leaf{},
		},
	}
	expected = "Composite with children: Leaf"
	if result := composite.Operation(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case: Composite with multiple children
	composite = &Composite{
		children: []Component{
			&Leaf{},
			&Composite{
				children: []Component{
					&Leaf{},
					&Leaf{},
				},
			},
			&Leaf{},
		},
	}
	expected = "Composite with children: Leaf Composite with children: Leaf Leaf Leaf "
	if result := composite.Operation(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
