package couponEvaluator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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
func Evaluate(requiredKeys []string, values map[string]interface{}, condition string) bool {
	loadConfiguration()
	bodyToEvaluate := MakeBody(requiredKeys, values, condition)
	headers := MakeHeaders()
	evaluateEndpoint := "/evaluate"
	response := doPost(configuration.Server+evaluateEndpoint, headers, bodyToEvaluate)
	if response == nil {
		return false
	}
	return ConvertToBoolean(response)
}

func ConvertToBoolean(data []byte) bool {
	boleanToReturn, err := strconv.ParseBool(string(data))
	valid := (err == nil)
	if valid {
		return boleanToReturn
	} else {
		log.Println("Invalid response")
		return false
	}
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
	request, _ := http.NewRequest("POST", url, body)
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	c1 := make(chan []byte, 1)
	go func() {
		response, err := client.Do(request)
		if err != nil {
			log.Println(err)
			c1 <- nil
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println(err)
			}
			if response.StatusCode != http.StatusOK {
				log.Println(response.StatusCode)
				hdr := response.Header
				for key, value := range hdr {
					log.Println("   ", key, ":", value)
				}
			}
			c1 <- contents
		}
	}()
	select {
	case result := <-c1:
		return result
	case <-time.After(10000 * time.Millisecond):
		return nil
	}
}
