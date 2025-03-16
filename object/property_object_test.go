package object_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dracory/base/object"
)

func Test_PropertyObject(t *testing.T) {
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

func Test_PropertyObject_Count(t *testing.T) {
	po := object.NewPropertyObject()

	// Initial count should be 0
	if count := po.Count(); count != 0 {
		t.Errorf("Initial count should be 0, got %d", count)
	}

	// Add one property
	po.Set("key1", "value1")
	if count := po.Count(); count != 1 {
		t.Errorf("Count should be 1 after adding one property, got %d", count)
	}

	// Add multiple properties
	po.Set("key2", 42)
	po.Set("key3", true)
	if count := po.Count(); count != 3 {
		t.Errorf("Count should be 3 after adding three properties, got %d", count)
	}

	// Overwrite existing property
	po.Set("key1", "new value")
	if count := po.Count(); count != 3 {
		t.Errorf("Count should still be 3 after overwriting a property, got %d", count)
	}

	// Remove a property
	po.Unset("key1")
	if count := po.Count(); count != 2 {
		t.Errorf("Count should be 2 after removing a property, got %d", count)
	}

	// Clear all properties
	po.Clear()
	if count := po.Count(); count != 0 {
		t.Errorf("Count should be 0 after clearing all properties, got %d", count)
	}
}

func Test_PropertyObject_Clear(t *testing.T) {
	po := object.NewPropertyObject()

	// Add several properties
	po.Set("key1", "value1")
	po.Set("key2", 42)
	po.Set("key3", true)

	// Verify properties exist
	if count := po.Count(); count != 3 {
		t.Errorf("Expected 3 properties before clear, got %d", count)
	}

	// Clear all properties
	po.Clear()

	// Verify count is 0
	if count := po.Count(); count != 0 {
		t.Errorf("Count should be 0 after Clear(), got %d", count)
	}

	// Verify no properties exist
	if po.Has("key1") {
		t.Error("Property 'key1' should not exist after Clear()")
	}
	if po.Has("key2") {
		t.Error("Property 'key2' should not exist after Clear()")
	}
	if po.Has("key3") {
		t.Error("Property 'key3' should not exist after Clear()")
	}

	// Verify getting cleared properties returns nil
	if val := po.Get("key1"); val != nil {
		t.Errorf("Get() should return nil for cleared property, got %v", val)
	}
}

func Test_PropertyObject_Get(t *testing.T) {
	po := object.NewPropertyObject()

	// Test getting non-existent property
	if val := po.Get("nonexistent"); val != nil {
		t.Errorf("Get() should return nil for non-existent property, got %v", val)
	}

	// Test with various data types
	testCases := map[string]interface{}{
		"string":       "test string",
		"empty_string": "",
		"int":          42,
		"zero_int":     0,
		"negative_int": -10,
		"float":        3.14159,
		"zero_float":   0.0,
		"bool_true":    true,
		"bool_false":   false,
		"nil_value":    nil,
		"slice":        []string{"a", "b", "c"},
		"empty_slice":  []string{},
		"map":          map[string]int{"a": 1, "b": 2},
		"empty_map":    map[string]int{},
		"struct":       struct{ Name string }{"test"},
	}

	// Set and get each test case
	for key, expected := range testCases {
		po.Set(key, expected)
		actual := po.Get(key)

		if expected == nil {
			if actual != nil {
				t.Errorf("Get(%q) should return nil, got %v", key, actual)
			}
		} else if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Get(%q) = %v, want %v", key, actual, expected)
		}
	}
}

func Test_PropertyObject_Has(t *testing.T) {
	po := object.NewPropertyObject()

	// Test Has on empty object
	if po.Has("any") {
		t.Error("Has() should return false for any key on empty object")
	}

	// Set some properties
	po.Set("existing", "value")
	po.Set("zero", 0)
	po.Set("empty", "")
	po.Set("nil", nil)

	// Test Has with existing properties
	if !po.Has("existing") {
		t.Error("Has() should return true for existing property")
	}
	if !po.Has("zero") {
		t.Error("Has() should return true for property with zero value")
	}
	if !po.Has("empty") {
		t.Error("Has() should return true for property with empty string")
	}
	if !po.Has("nil") {
		t.Error("Has() should return true for property with nil value")
	}

	// Test Has with non-existent property
	if po.Has("nonexistent") {
		t.Error("Has() should return false for non-existent property")
	}

	// Test after unsetting
	po.Unset("existing")
	if po.Has("existing") {
		t.Error("Has() should return false after Unset()")
	}

	// Test after clearing
	po.Clear()
	if po.Has("zero") || po.Has("empty") || po.Has("nil") {
		t.Error("Has() should return false for all properties after Clear()")
	}
}

func Test_PropertyObject_Keys(t *testing.T) {
	po := object.NewPropertyObject()

	// Test Keys on empty object
	keys := po.Keys()
	if len(keys) != 0 {
		t.Errorf("Keys() should return empty slice for empty object, got %v", keys)
	}

	// Add properties
	expectedKeys := []string{"key1", "key2", "key3"}
	po.Set("key1", "value1")
	po.Set("key2", 42)
	po.Set("key3", true)

	// Get keys and verify
	keys = po.Keys()
	if len(keys) != len(expectedKeys) {
		t.Errorf("Keys() returned %d keys, expected %d", len(keys), len(expectedKeys))
	}

	// Check that all expected keys are present
	// Note: Keys() doesn't guarantee order, so we need to check membership
	for _, expected := range expectedKeys {
		found := false
		for _, actual := range keys {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Keys() should contain %q", expected)
		}
	}

	// Test after removing a key
	po.Unset("key2")
	keys = po.Keys()
	if len(keys) != 2 {
		t.Errorf("Keys() should return 2 keys after Unset(), got %d", len(keys))
	}
	for _, key := range keys {
		if key == "key2" {
			t.Error("Keys() should not contain removed key")
		}
	}

	// Test after clearing
	po.Clear()
	keys = po.Keys()
	if len(keys) != 0 {
		t.Errorf("Keys() should return empty slice after Clear(), got %v", keys)
	}
}

func Test_PropertyObject_Set(t *testing.T) {
	po := object.NewPropertyObject()

	// Test basic set
	err := po.Set("key", "value")
	if err != nil {
		t.Errorf("Set() returned unexpected error: %v", err)
	}
	if val := po.Get("key"); val != "value" {
		t.Errorf("Set() failed to set value correctly, got %v", val)
	}

	// Test overwriting existing value
	err = po.Set("key", "new value")
	if err != nil {
		t.Errorf("Set() returned unexpected error when overwriting: %v", err)
	}
	if val := po.Get("key"); val != "new value" {
		t.Errorf("Set() failed to overwrite value correctly, got %v", val)
	}

	// Test with empty key
	err = po.Set("", "empty key")
	if err != nil {
		t.Errorf("Set() with empty key returned unexpected error: %v", err)
	}
	if val := po.Get(""); val != "empty key" {
		t.Errorf("Set() with empty key failed, got %v", val)
	}

	// Test with nil value
	err = po.Set("nil_key", nil)
	if err != nil {
		t.Errorf("Set() with nil value returned unexpected error: %v", err)
	}
	if po.Get("nil_key") != nil {
		t.Error("Set() with nil value failed")
	}
}

func Test_PropertyObject_Unset(t *testing.T) {
	po := object.NewPropertyObject()

	// Set some properties
	po.Set("key1", "value1")
	po.Set("key2", 42)

	// Unset existing property
	po.Unset("key1")
	if po.Has("key1") {
		t.Error("Unset() failed to remove property")
	}
	if val := po.Get("key1"); val != nil {
		t.Errorf("Get() after Unset() should return nil, got %v", val)
	}

	// Verify other properties are unaffected
	if !po.Has("key2") {
		t.Error("Unset() should not affect other properties")
	}

	// Unset non-existent property (should not cause error)
	po.Unset("nonexistent")

	// Unset with empty key
	po.Set("", "empty key")
	po.Unset("")
	if po.Has("") {
		t.Error("Unset() failed to remove property with empty key")
	}

	// Unset all properties
	po.Set("key3", true)
	po.Set("key4", 3.14)

	keys := po.Keys()
	for _, key := range keys {
		po.Unset(key)
	}

	if po.Count() != 0 {
		t.Errorf("Count should be 0 after unsetting all properties, got %d", po.Count())
	}
}

func Test_PropertyObject_WithComplexTypes(t *testing.T) {
	po := object.NewPropertyObject()

	// Test with nested structures
	type Person struct {
		Name    string
		Age     int
		Address struct {
			Street  string
			City    string
			Country string
		}
	}

	person := Person{
		Name: "John Doe",
		Age:  30,
	}
	person.Address.Street = "123 Main St"
	person.Address.City = "Anytown"
	person.Address.Country = "USA"

	// Set complex struct
	po.Set("person", person)

	// Get and verify
	retrieved := po.Get("person")
	if !reflect.DeepEqual(retrieved, person) {
		t.Errorf("Complex struct not stored/retrieved correctly, got %v, want %v", retrieved, person)
	}

	// Test with nested maps
	nestedMap := map[string]interface{}{
		"name": "Jane Doe",
		"details": map[string]interface{}{
			"age":      28,
			"employed": true,
			"skills":   []string{"Go", "Python", "JavaScript"},
		},
	}

	po.Set("nested_map", nestedMap)

	// Get and verify
	retrievedMap := po.Get("nested_map")
	if !reflect.DeepEqual(retrievedMap, nestedMap) {
		t.Errorf("Nested map not stored/retrieved correctly, got %v, want %v", retrievedMap, nestedMap)
	}
}

func Test_PropertyObject_Errors(t *testing.T) {
	// Create a new property object
	po := object.NewPropertyObject()

	// Test error case - getting a non-existent property
	propertyValue := po.Get("nonexistent")
	if propertyValue != nil {
		t.Error("Get should return nil for non-existent property")
	}
}

func Test_PropertyObject_ConcurrentOperations(t *testing.T) {
	po := object.NewPropertyObject()

	// Test concurrent operations
	// This test verifies that the PropertyObject can handle concurrent operations
	// safely with its mutex protection

	const numGoroutines = 10
	const numOperations = 100

	done := make(chan bool, numGoroutines)

	for i := range numGoroutines {
		go func(id int) {
			for j := 0; j < numOperations; j++ {
				key := "key" + string(rune('A'+id)) + string(rune('0'+j%10))

				// Set, Get, Has, Unset operations
				po.Set(key, j)
				po.Get(key)
				po.Has(key)

				if j%2 == 0 {
					po.Unset(key)
				}

				// Occasionally call Keys() and Count()
				if j%10 == 0 {
					po.Keys()
					po.Count()
				}

				// Occasionally Clear all properties
				if j%50 == 49 {
					po.Clear()
				}
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for range numGoroutines {
		<-done
	}

	// If we got here without panics or race detector issues, the test passes
}

func Test_PropertyObject_EdgeCases(t *testing.T) {
	po := object.NewPropertyObject()

	// Test with very long key
	longKey := ""
	for i := 0; i < 1000; i++ {
		longKey += "a"
	}
	po.Set(longKey, "long key value")
	if val := po.Get(longKey); val != "long key value" {
		t.Errorf("Failed with very long key, got %v", val)
	}

	// Test with special characters in key
	specialKeys := []string{
		"key with spaces",
		"key-with-hyphens",
		"key_with_underscores",
		"key.with.dots",
		"!@#$%^&*()",
		"日本語",     // Japanese
		"中文",      // Chinese
		"한국어",     // Korean
		"Русский", // Russian
	}

	for _, key := range specialKeys {
		po.Set(key, "special key value")
		if val := po.Get(key); val != "special key value" {
			t.Errorf("Failed with special key %q, got %v", key, val)
		}
	}

	// Test with large number of properties
	po.Clear()
	expectedCount := 1000
	for i := 0; i < expectedCount; i++ {
		po.Set(fmt.Sprintf("largeTest_%d", i), i)
	}

	// Get the actual count
	actualCount := po.Count()

	// Since we're using a thread-safe implementation, we should get exactly what we expect
	if actualCount != expectedCount {
		t.Errorf("Expected %d properties, got %d", expectedCount, actualCount)
	}

	keys := po.Keys()
	if len(keys) != expectedCount {
		t.Errorf("Expected %d keys, got %d", expectedCount, len(keys))
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
	retrievedValue := po.Get(key)
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
		actual := po.Get(k)
		if actual == nil {
			t.Fatalf("Get should not return nil for key '%s'", k)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Property %s value %v should match what was set %v", k, actual, expected)
		}
	}
}
