package config

import (
	"encoding/json"
	"errors"

	"github.com/dracory/base/object"
)

// Config represents the application configuration
type ConfigInterface interface {
	object.PropertyObjectInterface
	object.SerializableInterface
}

// config represents the application configuration
// implementation of the ConfigInterface
type configImplementation struct {
	propertyObj object.PropertyObjectInterface
	id          string
}

func NewConfig() ConfigInterface {
	return &configImplementation{
		propertyObj: object.NewPropertyObject(),
		id:          "",
	}
}

// Get implements PropertyObjectInterface
func (c *configImplementation) Get(key string) (interface{}, error) {
	return c.propertyObj.Get(key)
}

// Set implements PropertyObjectInterface
func (c *configImplementation) Set(key string, value interface{}) error {
	return c.propertyObj.Set(key, value)
}

// GetID implements SerializableInterface
func (c *configImplementation) GetID() string {
	if c.id != "" {
		return c.id
	}

	// Try to get ID from underlying property object if it's a PropertyObject
	if po, ok := c.propertyObj.(*object.PropertyObject); ok {
		return po.GetID()
	}

	return ""
}

// SetID implements SerializableInterface
func (c *configImplementation) SetID(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	c.id = id

	// Also set ID in underlying property object if it's a PropertyObject
	if po, ok := c.propertyObj.(*object.PropertyObject); ok {
		return po.SetID(id)
	}

	return nil
}

// ToJSON implements SerializableInterface
func (c *configImplementation) ToJSON() ([]byte, error) {
	// Try to use underlying property object's ToJSON if it's a PropertyObject
	if po, ok := c.propertyObj.(*object.PropertyObject); ok {
		return po.ToJSON()
	}

	// Fallback implementation
	properties := make(map[string]interface{})

	// Get all properties from the property object
	// This is a simplified implementation since we don't have a GetAll method
	// In a real implementation, you would iterate through all properties

	return json.Marshal(map[string]interface{}{
		"id":         c.GetID(),
		"properties": properties,
	})
}

// FromJSON implements SerializableInterface
func (c *configImplementation) FromJSON(data []byte) error {
	// Try to use underlying property object's FromJSON if it's a PropertyObject
	if po, ok := c.propertyObj.(*object.PropertyObject); ok {
		return po.FromJSON(data)
	}

	// Fallback implementation
	temp := struct {
		ID         string                 `json:"id"`
		Properties map[string]interface{} `json:"properties"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Set ID
	if temp.ID != "" {
		c.id = temp.ID
	}

	// Set properties
	if temp.Properties != nil {
		for k, v := range temp.Properties {
			if err := c.propertyObj.Set(k, v); err != nil {
				return err
			}
		}
	}

	return nil
}

// Ensure configImplementation implements ConfigInterface
var _ ConfigInterface = (*configImplementation)(nil)
