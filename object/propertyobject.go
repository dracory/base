package object

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type PropertyObjectInterface interface {
	Get(key string) (any, error)
	Set(key string, value any) error
	// Delete(key string) error
	// Has(key string) (bool, error)
	// Clear() error
	// Count() (int, error)
	// Keys() ([]string, error)
	// Values() ([]interface{}, error)
	// Items() ([][2]interface{}, error)
}

type PropertyObject struct {
	id         string
	properties map[string]any
}

func NewPropertyObject() PropertyObjectInterface {
	return &PropertyObject{
		id:         uuid.New().String(),
		properties: make(map[string]any),
	}
}

func (p *PropertyObject) GetID() string {
	return p.id
}

func (p *PropertyObject) SetID(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	p.id = id
	return nil
}

func (p *PropertyObject) Get(key string) (any, error) {
	value, exists := p.properties[key]
	if !exists {
		return nil, errors.New("property not found")
	}
	return value, nil
}

func (p *PropertyObject) Set(key string, value any) error {
	p.properties[key] = value
	return nil
}

// ToJSON serializes the PropertyObject to JSON
func (p *PropertyObject) ToJSON() ([]byte, error) {
	// Create a map that includes both the ID and properties
	data := map[string]interface{}{
		"id":         p.id,
		"properties": p.properties,
	}
	return json.Marshal(data)
}

// FromJSON deserializes JSON data into the PropertyObject
func (p *PropertyObject) FromJSON(data []byte) error {
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
		p.id = temp.ID
	}

	if temp.Properties != nil {
		p.properties = temp.Properties
	}

	return nil
}
