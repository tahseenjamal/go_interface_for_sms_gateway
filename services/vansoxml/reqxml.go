package vansoxml

import (
	"encoding/xml"
)

type ReqOperation struct {
	XMLName xml.Name `xml:"operation"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
	Account struct {
		Text     string `xml:",chardata"`
		Username string `xml:"username,attr"`
		Password string `xml:"password,attr"`
	} `xml:"account"`
	SubmitRequest struct {
		Chardata       string `xml:",chardata"`
		DeliveryReport string `xml:"deliveryReport"`
		SourceAddress  struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"sourceAddress"`
		DestinationAddress struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"destinationAddress"`
		Text struct {
			Text     string `xml:",chardata"`
			Encoding string `xml:"encoding,attr"`
		} `xml:"text"`
	} `xml:"submitRequest"`
}
