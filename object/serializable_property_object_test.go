package object_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/dracory/base/object"
	"github.com/dracory/base/testutils"
)

func TestSerializablePropertyObject(t *testing.T) {
	// Setup test environment
	testCfg := testutils.DefaultTestConfig()
	testutils.SetupTestEnvironment(testCfg)
	defer testutils.CleanupTestEnvironment(testCfg)

	// Create a new serializable property object
	spo := object.NewSerializablePropertyObject()

	// Test that the object is not nil
	if spo == nil {
		t.Fatal("NewSerializablePropertyObject should return a non-nil SerializablePropertyObjectInterface")
	}

	// Test that the object implements the SerializablePropertyObjectInterface
	_, ok := spo.(object.SerializablePropertyObjectInterface)
	if !ok {
		t.Fatal("NewSerializablePropertyObject should return an instance of SerializablePropertyObjectInterface")
	}

	// Test property operations (inherited from PropertyObject)
	testPropertyOperations(t, spo)

	// Test ID functionality
	testIDOperations(t, spo)

	// Test serialization
	testSerialization(t, spo)
}

func TestSerializablePropertyObjectErrors(t *testing.T) {
	// Setup test environment
	testCfg := testutils.DefaultTestConfig()
	testutils.SetupTestEnvironment(testCfg)
	defer testutils.CleanupTestEnvironment(testCfg)

	// Create a new serializable property object
	spo := object.NewSerializablePropertyObject()

	// Test error case - setting an empty ID
	spo.SetID("")
	if spo.GetID() == "" {
		t.Error("SetID should not allow empty ID")
	}

	// Test error case - getting a non-existent property
	propertyValue := spo.Get("nonexistent")
	if propertyValue != nil {
		t.Error("Get should return nil for non-existent property")
	}

	// Test error case - deserializing invalid JSON
	err := spo.FromJSON([]byte("invalid json"))
	if err == nil {
		t.Error("FromJSON should return an error for invalid JSON")
	}
}

// Helper function to test ID operations
func testIDOperations(t *testing.T, spo object.SerializablePropertyObjectInterface) {
	t.Helper()

	// Get the initial ID (should be auto-generated)
	initialID := spo.GetID()
	if initialID == "" {
		t.Error("Initial ID should not be empty")
	}

	// Set a new ID
	newID := "test-id-123"
	spo.SetID(newID)

	// Get the ID and verify it matches what was set
	retrievedID := spo.GetID()
	if retrievedID != newID {
		t.Errorf("Retrieved ID %s should match what was set %s", retrievedID, newID)
	}
}

// Helper function to test serialization
func testSerialization(t *testing.T, spo object.SerializablePropertyObjectInterface) {
	t.Helper()

	// Set some properties
	spo.Set("name", "Test Object")
	spo.Set("version", 1)
	spo.Set("enabled", true)
	spo.Set("tags", []string{"test", "serialization"})
	spo.Set("config", map[string]interface{}{
		"key1": "value1",
		"key2": 2,
	})

	// Set a specific ID for testing
	testID := "serialization-test-id"
	spo.SetID(testID)

	// Convert to JSON
	jsonData, err := spo.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON should not return an error: %v", err)
	}
	if len(jsonData) == 0 {
		t.Fatal("JSON data should not be empty")
	}

	// For debugging, print the JSON data
	t.Logf("JSON data: %s", string(jsonData))

	// Create a new serializable property object
	newSpo := object.NewSerializablePropertyObject()

	// Load from JSON
	err = newSpo.FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON should not return an error: %v", err)
	}

	// Verify ID was preserved
	if newSpo.GetID() != testID {
		t.Errorf("ID should be preserved during serialization, got %s, expected %s", newSpo.GetID(), testID)
	}

	// Verify simple properties were preserved
	verifyProperty(t, newSpo, "name", "Test Object")
	verifyNumericProperty(t, newSpo, "version", 1)
	verifyProperty(t, newSpo, "enabled", true)

	// Verify slice property
	tags := newSpo.Get("tags")

	// JSON unmarshaling might convert []string to []interface{}
	switch s := tags.(type) {
	case []string:
		if !reflect.DeepEqual(s, []string{"test", "serialization"}) {
			t.Errorf("Property 'tags' should match original value, got %v", s)
		}
	case []interface{}:
		expected := []interface{}{"test", "serialization"}
		if !reflect.DeepEqual(s, expected) {
			t.Errorf("Property 'tags' should match original value, got %v, expected %v", s, expected)
		}
	default:
		t.Errorf("Property 'tags' should be a slice type, got %T", tags)
	}

	// Verify map property
	config := newSpo.Get("config")

	// Convert both to JSON for comparison since types might differ
	expectedJSON, _ := json.Marshal(map[string]interface{}{
		"key1": "value1",
		"key2": 2,
	})
	actualJSON, _ := json.Marshal(config)

	if string(expectedJSON) != string(actualJSON) {
		t.Errorf("Property 'config' should match original value, got %s, expected %s", string(actualJSON), string(expectedJSON))
	}
}

// Helper function to verify a property value
func verifyProperty(t *testing.T, obj object.PropertyObjectInterface, key string, expected interface{}) {
	t.Helper()
	actual := obj.Get(key)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Property '%s' should have correct value after deserialization, got %v, expected %v", key, actual, expected)
	}
}

// Helper function to verify a numeric property value with type flexibility
func verifyNumericProperty(t *testing.T, obj object.PropertyObjectInterface, key string, expected float64) {
	t.Helper()
	actual := obj.Get(key)

	switch v := actual.(type) {
	case int:
		if float64(v) != expected {
			t.Errorf("Property '%s' should have value %f after deserialization, got %d", key, expected, v)
		}
	case int64:
		if float64(v) != expected {
			t.Errorf("Property '%s' should have value %f after deserialization, got %d", key, expected, v)
		}
	case float64:
		if v != expected {
			t.Errorf("Property '%s' should have value %f after deserialization, got %f", key, expected, v)
		}
	default:
		t.Errorf("Property '%s' should be a numeric type, got %T", key, actual)
	}
}
