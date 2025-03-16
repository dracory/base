package config

import (
	"github.com/dracory/base/object"
)

// ConfigInterface represents the application configuration
type ConfigInterface interface {
	object.SerializablePropertyObjectInterface
}

// configImplementation is the implementation of the ConfigInterface
type configImplementation struct {
	object.SerializablePropertyObjectInterface
}

// NewConfig creates a new configuration object
func NewConfig() ConfigInterface {
	return &configImplementation{
		SerializablePropertyObjectInterface: object.NewSerializablePropertyObject(),
	}
}

// Ensure configImplementation implements ConfigInterface
var _ ConfigInterface = (*configImplementation)(nil)
