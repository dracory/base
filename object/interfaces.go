package object

// PropertyObjectInterface defines the interface for objects
// that can store and retrieve properties
type PropertyObjectInterface interface {
	Clear()
	Count() int
	Get(key string) (any, error)
	Has(key string) bool
	Keys() []string
	Set(key string, value any)
	Unset(key string)
}

// SerializableInterface defines an interface for objects
// that can be serialized/deserialized
type SerializableInterface interface {
	// GetID returns the unique identifier for this object
	GetID() string

	// SetID sets the unique identifier for this object
	SetID(id string) error

	// ToJSON serializes the object to JSON
	ToJSON() ([]byte, error)

	// FromJSON deserializes JSON data into the object
	FromJSON(data []byte) error
}

// SerializablePropertyObjectInterface combines PropertyObjectInterface and SerializableInterface
type SerializablePropertyObjectInterface interface {
	PropertyObjectInterface
	SerializableInterface
}
