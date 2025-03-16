package object

import (
	"encoding/json"

	"github.com/google/uuid"
)

// SerializablePropertyObject extends PropertyObject with ID and serialization capabilities
type SerializablePropertyObject struct {
	PropertyObject
}

// NewSerializablePropertyObject creates a new SerializablePropertyObject with a generated UUID
func NewSerializablePropertyObject() SerializablePropertyObjectInterface {
	return &SerializablePropertyObject{
		PropertyObject: PropertyObject{properties: map[string]any{
			"id": uuid.New().String(),
		}},
	}
}

// GetID returns the unique identifier for this object
func (s *SerializablePropertyObject) GetID() string {
	if s.Has("id") {
		return s.Get("id").(string)
	}
	return ""
}

// SetID sets the unique identifier for this object
func (s *SerializablePropertyObject) SetID(id string) {
	if id == "" {
		return
	}
	s.Set("id", id)
}

// ToJSON serializes the SerializablePropertyObject to JSON
func (s *SerializablePropertyObject) ToJSON() ([]byte, error) {
	return json.Marshal(s.properties)
}

// FromJSON deserializes JSON data into the SerializablePropertyObject
func (s *SerializablePropertyObject) FromJSON(data []byte) error {
	// Create a temporary structure to hold the deserialized data
	temp := map[string]any{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.properties = temp

	return nil
}
