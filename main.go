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

type EvaluationBody struct {
	RequiredKeys []string               `json:"required_keys"`
	Values       map[string]interface{} `json:"values"`
	Condition    string                 `json:"condition"`
}

// Structs
type Configuration struct {
	Server string
}

// Functions
func Evaluate(requiredKeys []string, values map[string]interface{}, condition string) string {
	loadConfiguration()
	bodyToEvaluate := MakeBody(requiredKeys, values, "total < amount")
	headers := MakeHeaders()
	evaluateEndpoint := "/evaluate"
	response := doPost(configuration.Server+evaluateEndpoint, headers, bodyToEvaluate)
	return string(response)
}

func loadConfiguration() {
	file, _ := os.Open("config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
	}
	configuration.Server = "http://localhost:8082"
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
