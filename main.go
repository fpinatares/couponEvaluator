package couponEvaluator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Global variables
var configuration Configuration

// Structs
type Configuration struct {
	Server string
}

type EvaluationBody struct {
	RequiredKeys []string               `json:"required_keys"`
	Values       map[string]interface{} `json:"values"`
	Condition    string                 `json:"condition"`
}

// Functions
func Evaluate(requiredKeys []string, values map[string]interface{}, condition string) string {
	loadConfiguration()
	bodyToEvaluate := MakeBody(requiredKeys, values, condition)
	headers := MakeHeaders()
	evaluateEndpoint := "/evaluate"
	response := doPost(configuration.Server+evaluateEndpoint, headers, bodyToEvaluate)
	return string(response)
}

func loadConfiguration() {
	configuration.Server = os.Getenv("COUPONS_SERVER")
}

func MakeHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	return headers
}

func MakeBody(requiredKeys []string, values map[string]interface{}, condition string) *bytes.Buffer {
	var bodyToEvaluate EvaluationBody
	bodyToEvaluate.RequiredKeys = requiredKeys
	bodyToEvaluate.Values = values
	bodyToEvaluate.Condition = condition
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyToEvaluate)
	return body
}

func doPost(url string, headers map[string]string, body *bytes.Buffer) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, body)
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		if response.StatusCode != http.StatusOK {
			fmt.Println(response.StatusCode)
			hdr := response.Header
			for key, value := range hdr {
				fmt.Println("   ", key, ":", value)
			}
		}
		return contents
	}
}
