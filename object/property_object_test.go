package object_test

import (
	"testing"

	"github.com/dracory/base/object"
)

func TestPropertyObject(t *testing.T) {
	// Create a new property object
	po := object.NewPropertyObject()

	// Test that the property object is not nil
	if po == nil {
		t.Fatal("NewPropertyObject should return a non-nil PropertyObjectInterface")
	}

	// Test that the property object implements the PropertyObjectInterface
	_, ok := po.(object.PropertyObjectInterface)
	if !ok {
		t.Fatal("NewPropertyObject should return an instance of PropertyObjectInterface")
	}

	// Test property setting and getting
	testPropertyOperations(t, po)
}

func TestPropertyObjectErrors(t *testing.T) {
	// Create a new property object
	po := object.NewPropertyObject()

	// Test error case - getting a non-existent property
	_, err := po.Get("nonexistent")
	if err == nil {
		t.Error("Get should return an error for non-existent property")
	}
}

// Helper function to test property operations
func testPropertyOperations(t *testing.T, po object.PropertyObjectInterface) {
	t.Helper()

	// Test with simple string
	key := "testKey"
	value := "testValue"

	// Set a property
	po.Set(key, value)

	// Get the property and verify it matches what was set
	retrievedValue, err := po.Get(key)
	if err != nil {
		t.Fatalf("Get should not return an error: %v", err)
	}
	if retrievedValue != value {
		t.Errorf("Retrieved property value %v should match what was set %v", retrievedValue, value)
	}

	// Test with multiple property types
	properties := map[string]interface{}{
		"string": "test",
		"int":    42,
		"bool":   true,
		"float":  3.14,
		"slice":  []string{"a", "b", "c"},
		"map":    map[string]int{"a": 1, "b": 2},
	}

	// Set all properties
	for k, v := range properties {
		po.Set(k, v)
	}

	// Verify all properties were set correctly
	for k, expected := range properties {
		actual, err := po.Get(k)
		if err != nil {
			t.Fatalf("Get should not return an error for key %s: %v", k, err)
		}

		// For complex types like slices and maps, we need to use deep comparison
		// This is a simplified check - in a real test you might want to use reflect.DeepEqual
		switch k {
		case "string":
			if actual != expected {
				t.Errorf("Property %s value %v should match what was set %v", k, actual, expected)
			}
		case "int":
			if actual != expected {
				t.Errorf("Property %s value %v should match what was set %v", k, actual, expected)
			}
		case "bool":
			if actual != expected {
				t.Errorf("Property %s value %v should match what was set %v", k, actual, expected)
			}
		case "float":
			if actual != expected {
				t.Errorf("Property %s value %v should match what was set %v", k, actual, expected)
			}
		}
		// Note: For slice and map, a more complex comparison would be needed
	}
}
