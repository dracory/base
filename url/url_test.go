package url

import (
	"testing"
)

func TestURLBuilder(t *testing.T) {
	// Test creating a new URLBuilder
	builder := NewURLBuilder("https://example.com")

	if builder.RootURL() != "https://example.com" {
		t.Errorf("Expected https://example.com, got %s", builder.RootURL())
	}

	// Test with empty root URL
	emptyBuilder := NewURLBuilder("")
	if emptyBuilder.RootURL() != "" {
		t.Errorf("Expected empty string, got %s", emptyBuilder.RootURL())
	}
}

func TestURLBuilder_BuildURL(t *testing.T) {
	builder := NewURLBuilder("https://example.com")

	// Test basic path
	result := builder.BuildURL("test/path", nil)
	expected := "https://example.com/test/path"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test path with leading slash
	result = builder.BuildURL("/test/path", nil)
	expected = "https://example.com/test/path"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test path with query parameters
	params := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	result = builder.BuildURL("test/path", params)
	expected = "https://example.com/test/path?key1=value1&key2=value2"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test empty path
	result = builder.BuildURL("", nil)
	expected = "https://example.com/"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test with empty root URL
	emptyBuilder := NewURLBuilder("")
	result = emptyBuilder.BuildURL("test/path", nil)
	expected = "/test/path"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestURLBuilder_BuildQuery(t *testing.T) {
	builder := NewURLBuilder("https://example.com")

	// Test empty map
	result := builder.BuildQuery(map[string]string{})
	expected := ""
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test nil map
	result = builder.BuildQuery(nil)
	expected = ""
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test single parameter
	params := map[string]string{"key": "value"}
	result = builder.BuildQuery(params)
	expected = "?key=value"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test multiple parameters
	params = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	result = builder.BuildQuery(params)
	expected = "?key1=value1&key2=value2"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test parameter with special characters
	params = map[string]string{"search": "hello world"}
	result = builder.BuildQuery(params)
	expected = "?search=hello+world"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestURLBuilder_HttpBuildQuery(t *testing.T) {
	builder := NewURLBuilder("https://example.com")

	// Test empty values
	values := make(map[string][]string)
	result := builder.HttpBuildQuery(values)
	expected := ""
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test single value
	values = map[string][]string{"key": {"value"}}
	result = builder.HttpBuildQuery(values)
	expected = "key=value"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test multiple values (should include all values)
	values = map[string][]string{"key": {"value1", "value2"}}
	result = builder.HttpBuildQuery(values)
	expected = "key=value1&key=value2"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSetDefaultURL(t *testing.T) {
	// Test setting default URL
	SetDefaultURL("https://default.com")

	if RootURL() != "https://default.com" {
		t.Errorf("Expected https://default.com, got %s", RootURL())
	}

	// Test BuildURL with default
	result := BuildURL("test", nil)
	expected := "https://default.com/test"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test with empty default URL
	SetDefaultURL("")
	if RootURL() != "" {
		t.Errorf("Expected empty string, got %s", RootURL())
	}

	result = BuildURL("test", nil)
	expected = "/test"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
