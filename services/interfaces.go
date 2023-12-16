package services

import (
	"net/http"
	"time"

	"github.com/magiconair/properties"
)

type Parameters interface {
	configure(string)
}

type Connection interface {
	messageProcess(string) string
	setConfiguration(properties.Properties)
	sendMessage(string, string, string) (string, error)
	GetConfiguration() interface{}
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type VansoConfig struct {
	Url              string        `properties:"url"`
	VansoHTTPtimeout time.Duration `properties:"httptimeout"`
	Type             string        `properties:"type"`
	Username         string        `properties:"username"`
	Password         string        `properties:"password"`
	DeliveryReport   string        `properties:"deliveryreport"`
	SrcType          string        `properties:"srctype"`
	SrcText          string        `properties:"srctext"`
	DesType          string        `properties:"destype"`
	Encoding         string        `properties:"encoding"`
}
