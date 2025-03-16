package config_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/dracory/base/config"
	"github.com/dracory/base/testutils"
)

func TestNewConfig(t *testing.T) {
	// Setup test environment
	testCfg := testutils.DefaultTestConfig()
	testutils.SetupTestEnvironment(testCfg)
	defer testutils.CleanupTestEnvironment(testCfg)

	// Create a new config instance
	cfg := config.NewConfig()

	// Test that the config instance is not nil
	if cfg == nil {
		t.Fatal("NewConfig should return a non-nil ConfigInterface")
	}

	// Test that the config instance implements the ConfigInterface
	_, ok := cfg.(config.ConfigInterface)
	if !ok {
		t.Fatal("NewConfig should return an instance of ConfigInterface")
	}

	// Test property setting and getting
	key := "testKey"
	value := "testValue"
	
	// Set a property
	err := cfg.Set(key, value)
	if err != nil {
		t.Fatalf("Set should not return an error: %v", err)
	}
	
	// Get the property and verify it matches what was set
	retrievedValue, err := cfg.Get(key)
	if err != nil {
		t.Fatalf("Get should not return an error: %v", err)
	}
	if retrievedValue != value {
		t.Errorf("Retrieved property value %v should match what was set %v", retrievedValue, value)
	}

	// Test ID functionality (from Serializable interface)
	// Set an ID
	testID := "test-id-123"
	err = cfg.SetID(testID)
	if err != nil {
		t.Fatalf("SetID should not return an error: %v", err)
	}
	
	// Get the ID and verify it matches
	retrievedID := cfg.GetID()
	if retrievedID != testID {
		t.Errorf("Retrieved ID %s should match what was set %s", retrievedID, testID)
	}

	// Test error case - getting a non-existent property
	_, err = cfg.Get("nonexistent")
	if err == nil {
		t.Error("Get should return an error for non-existent property")
	}
}

func TestConfigProperties(t *testing.T) {
	// Setup test environment
	testCfg := testutils.DefaultTestConfig()
	testutils.SetupTestEnvironment(testCfg)
	defer testutils.CleanupTestEnvironment(testCfg)

	// Create a new config instance
	cfg := config.NewConfig()

	// Test multiple properties
	properties := map[string]interface{}{
		"string":  "test",
		"int":     42,
		"bool":    true,
		"float":   3.14,
		"slice":   []string{"a", "b", "c"},
		"map":     map[string]int{"a": 1, "b": 2},
	}

	// Set all properties
	for k, v := range properties {
		err := cfg.Set(k, v)
		if err != nil {
			t.Fatalf("Set should not return an error for key %s: %v", k, err)
		}
	}

	// Verify all properties were set correctly
	for k, expected := range properties {
		actual, err := cfg.Get(k)
		if err != nil {
			t.Fatalf("Get should not return an error for key %s: %v", k, err)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Property %s value %v should match what was set %v", k, actual, expected)
		}
	}

	// Test error case - getting a non-existent property
	_, err := cfg.Get("nonexistent")
	if err == nil {
		t.Error("Get should return an error for non-existent property")
	}
}

func TestConfigSerialization(t *testing.T) {
	// Setup test environment
	testCfg := testutils.DefaultTestConfig()
	testutils.SetupTestEnvironment(testCfg)
	defer testutils.CleanupTestEnvironment(testCfg)

	// Create a new config instance
	cfg := config.NewConfig()

	// Set some properties
	cfg.Set("name", "Test Config")
	cfg.Set("version", 1)
	cfg.Set("enabled", true)

	// Set an ID
	testID := "config-123"
	cfg.SetID(testID)

	// Convert to JSON
	jsonData, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON should not return an error: %v", err)
	}
	if len(jsonData) == 0 {
		t.Fatal("JSON data should not be empty")
	}

	// Create a new config instance
	newCfg := config.NewConfig()

	// Load from JSON
	err = newCfg.FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON should not return an error: %v", err)
	}

	// Verify ID was preserved
	if newCfg.GetID() != testID {
		t.Errorf("ID should be preserved during serialization, got %s, expected %s", newCfg.GetID(), testID)
	}

	// Verify properties were preserved
	name, err := newCfg.Get("name")
	if err != nil {
		t.Fatalf("Get should not return an error for key 'name': %v", err)
	}
	if name != "Test Config" {
		t.Errorf("Property 'name' should have correct value after deserialization, got %v, expected %s", name, "Test Config")
	}

	// For numeric values, we need to be more flexible due to JSON type conversions
	version, err := newCfg.Get("version")
	if err != nil {
		t.Fatalf("Get should not return an error for key 'version': %v", err)
	}
	
	// Check if the version is a number with value 1, regardless of specific numeric type
	switch v := version.(type) {
	case int:
		if v != 1 {
			t.Errorf("Property 'version' should have value 1 after deserialization, got %d", v)
		}
	case int64:
		if v != 1 {
			t.Errorf("Property 'version' should have value 1 after deserialization, got %d", v)
		}
	case float64:
		if v != 1.0 {
			t.Errorf("Property 'version' should have value 1.0 after deserialization, got %f", v)
		}
	default:
		t.Errorf("Property 'version' should be a numeric type, got %T", version)
	}

	enabled, err := newCfg.Get("enabled")
	if err != nil {
		t.Fatalf("Get should not return an error for key 'enabled': %v", err)
	}
	if enabled != true {
		t.Errorf("Property 'enabled' should have correct value after deserialization, got %v, expected %v", enabled, true)
	}
}

// TestJSONRoundTrip tests that a config object can be serialized to JSON and back
func TestJSONRoundTrip(t *testing.T) {
	// Setup test environment
	testCfg := testutils.DefaultTestConfig()
	testutils.SetupTestEnvironment(testCfg)
	defer testutils.CleanupTestEnvironment(testCfg)

	// Create a new config instance with various property types
	cfg := config.NewConfig()
	cfg.Set("string", "test")
	cfg.Set("int", 42)
	cfg.Set("bool", true)
	cfg.Set("float", 3.14)
	cfg.Set("slice", []string{"a", "b", "c"})
	cfg.Set("nested", map[string]interface{}{
		"key1": "value1",
		"key2": 2,
	})

	// Set an ID
	testID := "config-round-trip-123"
	cfg.SetID(testID)

	// Convert to JSON
	jsonData, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON should not return an error: %v", err)
	}

	// For debugging, print the JSON data
	t.Logf("JSON data: %s", string(jsonData))

	// Create a new config instance
	newCfg := config.NewConfig()

	// Load from JSON
	err = newCfg.FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON should not return an error: %v", err)
	}

	// Verify ID was preserved
	if newCfg.GetID() != testID {
		t.Errorf("ID should be preserved during serialization, got %s, expected %s", newCfg.GetID(), testID)
	}

	// Verify we can get all properties back (with type flexibility for numbers)
	verifyProperty(t, newCfg, "string", "test")
	verifyNumericProperty(t, newCfg, "int", 42)
	verifyProperty(t, newCfg, "bool", true)
	verifyNumericProperty(t, newCfg, "float", 3.14)
	
	// For slices and maps, we need to use reflect.DeepEqual
	slice, err := newCfg.Get("slice")
	if err != nil {
		t.Fatalf("Get should not return an error for key 'slice': %v", err)
	}
	
	// JSON unmarshaling might convert []string to []interface{}
	switch s := slice.(type) {
	case []string:
		if !reflect.DeepEqual(s, []string{"a", "b", "c"}) {
			t.Errorf("Property 'slice' should match original value, got %v", s)
		}
	case []interface{}:
		expected := []interface{}{"a", "b", "c"}
		if !reflect.DeepEqual(s, expected) {
			t.Errorf("Property 'slice' should match original value, got %v, expected %v", s, expected)
		}
	default:
		t.Errorf("Property 'slice' should be a slice type, got %T", slice)
	}

	// Check nested map
	nested, err := newCfg.Get("nested")
	if err != nil {
		t.Fatalf("Get should not return an error for key 'nested': %v", err)
	}
	
	// Convert both to JSON for comparison since types might differ
	expectedJSON, _ := json.Marshal(map[string]interface{}{
		"key1": "value1",
		"key2": 2,
	})
	actualJSON, _ := json.Marshal(nested)
	
	if string(expectedJSON) != string(actualJSON) {
		t.Errorf("Property 'nested' should match original value, got %s, expected %s", string(actualJSON), string(expectedJSON))
	}
}

// Helper function to verify a property value
func verifyProperty(t *testing.T, cfg config.ConfigInterface, key string, expected interface{}) {
	t.Helper()
	actual, err := cfg.Get(key)
	if err != nil {
		t.Fatalf("Get should not return an error for key '%s': %v", key, err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Property '%s' should have correct value after deserialization, got %v, expected %v", key, actual, expected)
	}
}

// Helper function to verify a numeric property value with type flexibility
func verifyNumericProperty(t *testing.T, cfg config.ConfigInterface, key string, expected float64) {
	t.Helper()
	actual, err := cfg.Get(key)
	if err != nil {
		t.Fatalf("Get should not return an error for key '%s': %v", key, err)
	}
	
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
