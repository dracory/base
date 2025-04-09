# Object Package

The `object` package provides a flexible and thread-safe implementation for managing properties and serializable objects in Go applications.

## Overview

This package offers a set of interfaces and implementations for:

1. **Property Objects**: Thread-safe containers for storing and retrieving key-value pairs
2. **Serializable Objects**: Objects that can be converted to and from JSON
3. **Combined Functionality**: Objects that are both property containers and serializable

## Interfaces

### PropertyObjectInterface

Defines methods for objects that can store and retrieve properties:

```go
type PropertyObjectInterface interface {
    Clear()                    // Removes all properties
    Count() int                // Returns the number of properties
    Get(key string) any        // Retrieves a property value by key
    Has(key string) bool       // Checks if a property exists
    Keys() []string            // Returns all property keys
    Set(key string, value any) error  // Stores a property value
    Unset(key string)          // Removes a property by key
}
```

### SerializableInterface

Defines methods for objects that can be serialized/deserialized:

```go
type SerializableInterface interface {
    GetID() string             // Returns the unique identifier
    SetID(id string)           // Sets the unique identifier
    ToJSON() ([]byte, error)   // Serializes the object to JSON
    FromJSON(data []byte) error // Deserializes JSON data into the object
}
```

### SerializablePropertyObjectInterface

Combines both interfaces for objects that are both property containers and serializable:

```go
type SerializablePropertyObjectInterface interface {
    PropertyObjectInterface
    SerializableInterface
}
```

## Implementations

### PropertyObject

A basic implementation of `PropertyObjectInterface` that provides thread-safe property storage:

```go
obj := object.NewPropertyObject()
obj.Set("name", "John")
obj.Set("age", 30)
name := obj.Get("name")  // "John"
keys := obj.Keys()       // ["name", "age"]
```

### SerializablePropertyObject

Extends `PropertyObject` with ID and serialization capabilities:

```go
obj := object.NewSerializablePropertyObject()
obj.Set("name", "John")
obj.Set("age", 30)

// Serialize to JSON
jsonData, err := obj.ToJSON()

// Deserialize from JSON
newObj := object.NewSerializablePropertyObject()
newObj.FromJSON(jsonData)
```

## Features

- **Thread-safe**: All operations are protected by mutex locks
- **Flexible**: Store any type of value (using Go's `any` type)
- **Serializable**: Convert objects to and from JSON
- **Unique IDs**: Automatic UUID generation for serializable objects
- **Extensible**: Easy to implement custom objects that satisfy the interfaces
- **Safe Property Access**: Built-in property existence checking with `Has` method

## Usage Examples

### Basic Property Object

```go
import "github.com/dracory/base/object"

// Create a new property object
obj := object.NewPropertyObject()

// Set properties
obj.Set("name", "John")
obj.Set("age", 30)

// Get properties
name := obj.Get("name")  // "John"
age := obj.Get("age")    // 30

// Check if property exists
hasName := obj.Has("name")  // true
hasEmail := obj.Has("email")  // false

// Get all keys
keys := obj.Keys()  // ["name", "age"]

// Remove a property
obj.Unset("age")

// Clear all properties
obj.Clear()
```

### Safe Property Access with Has

The `Has` method is crucial for safe property access. Here are some common patterns:

```go
obj := object.NewPropertyObject()

// Pattern 1: Check before getting
if obj.Has("email") {
    email := obj.Get("email").(string)
    // Use email safely
}

// Pattern 2: Type assertion with existence check
if obj.Has("age") {
    if age, ok := obj.Get("age").(int); ok {
        // Use age safely
    }
}

// Pattern 3: Conditional property access
if obj.Has("settings") && obj.Has("settings.notifications") {
    notifications := obj.Get("settings.notifications").(bool)
    // Use notifications safely
}

// Pattern 4: Default value pattern
var name string
if obj.Has("name") {
    name = obj.Get("name").(string)
} else {
    name = "Default Name"
}
```

### Serializable Property Object

```go
import "github.com/dracory/base/object"

// Create a new serializable property object
obj := object.NewSerializablePropertyObject()

// Set properties
obj.Set("name", "John")
obj.Set("age", 30)

// Get the object ID
id := obj.GetID()  // UUID string

// Set a custom ID
obj.SetID("custom-id")

// Serialize to JSON
jsonData, err := obj.ToJSON()
if err != nil {
    // Handle error
}

// Create a new object and deserialize
newObj := object.NewSerializablePropertyObject()
err = newObj.FromJSON(jsonData)
if err != nil {
    // Handle error
}

// Access properties safely
if newObj.Has("name") {
    name := newObj.Get("name").(string)
    // Use name safely
}
```

## Best Practices

1. **Always Check Property Existence**: Use `Has` before accessing properties to prevent nil pointer dereferences
2. **Type Assertions**: Combine `Has` with type assertions for type-safe property access
3. **Default Values**: Use `Has` to implement default value patterns
4. **Nested Properties**: Check existence of nested properties before access
5. **Thread Safety**: All methods including `Has` are thread-safe

## License

This package is part of the dracory/base project and is licensed under the same terms. 