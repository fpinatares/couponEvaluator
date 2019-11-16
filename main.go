package couponEvaluator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type EvaluationBody struct {
	RequiredKeys []string               `json:"required_keys"`
	Values       map[string]interface{} `json:"values"`
	Condition    string                 `json:"condition"`
}

func Evaluate(requiredKeys []string, values map[string]interface{}, condition string) { //(bool, error) {
	bodyToEvaluate := EvaluationBody{}
	bodyToEvaluate.RequiredKeys = requiredKeys
	bodyToEvaluate.Values = values
	bodyToEvaluate.Condition = condition
	/*	bodyToEvaluate, err := json.Marshal(map[string]string{
		"required_keys": requiredKeys,
		"values":        values,
		"condition":     condition,
	})*/
	bodyToRequest, err := json.Marshal(bodyToEvaluate)
	response, err := http.Post("localhost:8082/evaluate", "application/json", bytes.NewBuffer(bodyToRequest))
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
