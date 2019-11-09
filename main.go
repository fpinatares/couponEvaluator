package couponEvaluator

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func Evaluate(required_keys []string, values map[string]interface{}, condition string) bool {
	expression, err := govaluate.NewEvaluableExpression(condition)
	result, err := expression.Evaluate(values)
	fmt.Println(result)
	fmt.Println("ERROR ", err)
	return false
}
