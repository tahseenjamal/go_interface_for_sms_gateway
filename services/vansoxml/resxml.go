package vansoxml

import "encoding/xml"

type ResOperation struct {
	XMLName        xml.Name `xml:"operation"`
	Text           string   `xml:",chardata"`
	Type           string   `xml:"type,attr"`
	SubmitResponse struct {
		Text  string `xml:",chardata"`
		Error struct {
			Text    string `xml:",chardata"`
			Code    string `xml:"code,attr"`
			Message string `xml:"message,attr"`
		} `xml:"error"`
		TicketId string `xml:"ticketId"`
	} `xml:"submitResponse"`
}
