package services

import (
	"github.com/magiconair/properties"
)

type ConnParameters struct {
	configuration *properties.Properties
}

func NewConnParameters(filename string) ConnParameters {

	var temp ConnParameters

	temp.configure(filename)

	return temp
}

func (v *ConnParameters) configure(filename string) error {

	temp := properties.MustLoadFile("config.properties", properties.UTF8)

	v.configuration = temp

	return nil
}

func (v *ConnParameters) GetParameters() properties.Properties {

	return *v.configuration
}

func NewConnection(connHandler Connection, rawConfiguration properties.Properties) Connection {

	connHandler.setConfiguration(rawConfiguration)
	return connHandler
}
