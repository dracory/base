package object

import (
	"sync"
)

// PropertyObject is a basic implementation of PropertyObjectInterface
type PropertyObject struct {
	properties map[string]any
	mutex      sync.RWMutex
}

// NewPropertyObject creates a new PropertyObject
func NewPropertyObject() PropertyObjectInterface {
	return &PropertyObject{
		properties: make(map[string]any),
	}
}

// Count returns the number of properties
func (p *PropertyObject) Count() int {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return len(p.properties)
}

// Clear removes all properties
func (p *PropertyObject) Clear() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.properties = make(map[string]any)
}

// Get retrieves a property value by key
func (p *PropertyObject) Get(key string) any {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	value, exists := p.properties[key]
	if exists {
		return value
	}
	return nil
}

// Has checks if a property exists
func (p *PropertyObject) Has(key string) bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	_, exists := p.properties[key]
	return exists
}

// Keys returns all property keys
func (p *PropertyObject) Keys() []string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	keys := make([]string, 0, len(p.properties))
	for k := range p.properties {
		keys = append(keys, k)
	}
	return keys
}

// Set stores a property value with the given key
func (p *PropertyObject) Set(key string, value any) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.properties[key] = value
	return nil
}

// Unset removes a property by key
func (p *PropertyObject) Unset(key string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.properties, key)
}
