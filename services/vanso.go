package services

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sms/services/vansoxml"
	"strings"

	"github.com/magiconair/properties"
)

type VansoConnection struct {
	configurations VansoConfig
}

func (v *VansoConnection) setConfiguration(rawConfiguration properties.Properties) {

	if err := rawConfiguration.Decode(&v.configurations); err != nil {
		log.Fatal(err)
	}

}

func (v *VansoConnection) GetConfiguration() interface{} {

	return v.configurations
}

func (v *VansoConnection) messageProcess(message string) string {

	return hex.EncodeToString([]byte(message))

}

func (v *VansoConnection) reqXMLBody(mobile string, message string) string {

	var messageHex string = v.messageProcess(message)

	var operation vansoxml.ReqOperation

	operation.Type = v.configurations.Type
	operation.Account.Username = v.configurations.Username
	operation.Account.Password = v.configurations.Password
	operation.SubmitRequest.DeliveryReport = v.configurations.DeliveryReport
	operation.SubmitRequest.SourceAddress.Type = v.configurations.SrcType
	operation.SubmitRequest.SourceAddress.Text = v.configurations.SrcText
	operation.SubmitRequest.DestinationAddress.Type = v.configurations.DesType
	operation.SubmitRequest.Text.Encoding = v.configurations.Encoding
	operation.SubmitRequest.DestinationAddress.Text = mobile
	operation.SubmitRequest.Text.Text = messageHex

	out, _ := xml.MarshalIndent(operation, "", " ")
	var xmlbody string = xml.Header + string(out)

	return xmlbody

}
func (v *VansoConnection) sendMessage(mobile string, message string, test string) (string, error) {

	var res string

	if test == "true" {

		res = "OK"

	} else {
		res = v.resXMLParse(v.callHTTPClient(v.reqXMLBody(mobile, message)))
	}

	return res, nil

}

func (v *VansoConnection) callHTTPClient(xmlbody string) []byte {

	client := http.Client{
		Timeout: v.configurations.VansoHTTPtimeout, //* time.Millisecond,
	}

	var url string = v.configurations.Url
	resp, err := client.Post(url, "text/xml", strings.NewReader(xmlbody))

	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var bodyBytes []byte

	if resp.StatusCode == http.StatusOK {
		var err error
		bodyBytes, err = io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(bodyBytes))

		return bodyBytes

	}

	return bodyBytes
}

func (v *VansoConnection) resXMLParse(res []byte) string {

	var operation vansoxml.ResOperation

	xml.Unmarshal(res, &operation)

	return operation.SubmitResponse.Error.Message

}

func (v *VansoConnection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	q, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}

	var mobile string = q.Get("dest")
	var message string = q.Get("message")
	var test string = q.Get("test")
	res, _ := v.sendMessage(mobile, message, test)

	fmt.Fprintf(w, "001 %s", res)
}
