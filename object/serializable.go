package object

// Serializable defines an interface for objects that can be serialized/deserialized
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

type SerializablePropertyObjectInterface interface {
	PropertyObjectInterface
	SerializableInterface
}
