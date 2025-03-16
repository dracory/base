package object

import (
	"errors"
)

// PropertyObject is a basic implementation of PropertyObjectInterface
type PropertyObject struct {
	properties map[string]any
}

// NewPropertyObject creates a new PropertyObject
func NewPropertyObject() PropertyObjectInterface {
	return &PropertyObject{
		properties: make(map[string]any),
	}
}

// Count returns the number of properties
func (p *PropertyObject) Count() int {
	return len(p.properties)
}

// Clear removes all properties
func (p *PropertyObject) Clear() {
	p.properties = make(map[string]any)
}

// Get retrieves a property value by key
func (p *PropertyObject) Get(key string) (any, error) {
	if p.Has(key) {
		return p.properties[key], nil
	}
	return nil, errors.New("property not found")
}

// Has checks if a property exists
func (p *PropertyObject) Has(key string) bool {
	_, exists := p.properties[key]
	return exists
}

// Keys returns all property keys
func (p *PropertyObject) Keys() []string {
	keys := make([]string, 0, len(p.properties))
	for k := range p.properties {
		keys = append(keys, k)
	}
	return keys
}

// Set stores a property value with the given key
func (p *PropertyObject) Set(key string, value any) {
	p.properties[key] = value
}

// Unset removes a property by key
func (p *PropertyObject) Unset(key string) {
	delete(p.properties, key)
}
