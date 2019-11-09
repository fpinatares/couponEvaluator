package couponEvaluator

import (
	"github.com/Knetic/govaluate"
)

func Evaluate(required_keys []string, values map[string]interface{}, condition string) (bool, error) {
	if CheckValidValues(required_keys, values) {
		expression, err := govaluate.NewEvaluableExpression(condition)
		if err != nil {
			return false, err
		} else {
			result, err := expression.Evaluate(values)
			if err != nil {
				return false, err
			} else {
				return result.(bool), nil
			}
		}
	} else {
		return false, nil
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
