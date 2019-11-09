package couponEvaluator

import (
	"github.com/Knetic/govaluate"
)

func Evaluate(required_keys []string, values map[string]interface{}, condition string) bool {
	if CheckValidValues(required_keys, values) {
		expression, err := govaluate.NewEvaluableExpression(condition)
		result, err := expression.Evaluate(values)
		return err == nil && result.(bool)
	} else {
		return false
	}
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
