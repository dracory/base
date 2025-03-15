package config

import "github.com/gouniverse/dataobject"

// Config represents the application configuration
type ConfigInterface interface {
	dataobject.DataObjectInterface
}

// sonfig represents the application configuration
type configImplementation struct {
	dataobject.DataObject
}

func NewConfig() ConfigInterface {
	o := dataobject.NewDataObject()
	return &configImplementation{
		DataObject: *o,
	}
}
