package object

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// SerializablePropertyObject extends PropertyObject with ID and serialization capabilities
type SerializablePropertyObject struct {
	id string
	PropertyObject
}

// NewSerializablePropertyObject creates a new SerializablePropertyObject with a generated UUID
func NewSerializablePropertyObject() SerializablePropertyObjectInterface {
	return &SerializablePropertyObject{
		id:             uuid.New().String(),
		PropertyObject: PropertyObject{properties: make(map[string]any)},
	}
}

// GetID returns the unique identifier for this object
func (s *SerializablePropertyObject) GetID() string {
	return s.id
}

// SetID sets the unique identifier for this object
func (s *SerializablePropertyObject) SetID(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	s.id = id
	return nil
}

// ToJSON serializes the SerializablePropertyObject to JSON
func (s *SerializablePropertyObject) ToJSON() ([]byte, error) {
	// Create a map that includes both the ID and properties
	data := map[string]interface{}{
		"id":         s.id,
		"properties": s.properties,
	}
	return json.Marshal(data)
}

// FromJSON deserializes JSON data into the SerializablePropertyObject
func (s *SerializablePropertyObject) FromJSON(data []byte) error {
	// Create a temporary structure to hold the deserialized data
	temp := struct {
		ID         string         `json:"id"`
		Properties map[string]any `json:"properties"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Update the object with the deserialized data
	if temp.ID != "" {
		s.id = temp.ID
	}

	if temp.Properties != nil {
		s.properties = temp.Properties
	}

	return nil
}
