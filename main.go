package couponEvaluator

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func Evaluate(required_keys []string, values map[string]interface{}, condition string) bool {
	applies := false
	if CheckValidValues(required_keys, values) {
		expression, err := govaluate.NewEvaluableExpression(condition)
		result, err := expression.Evaluate(values)
		if err != nil {
			fmt.Println("ERROR: ", err)
		} else {
			applies = result.(bool)
		}
	}
	return applies
}

func CheckValidValues(required_keys []string, values map[string]interface{}) bool {
	valid := true
	for i := range required_keys {
		if _, ok := values[required_keys[i]]; !ok {
			valid = false
		}
	}
	return valid
}
