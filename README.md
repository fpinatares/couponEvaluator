# couponEvaluator
This repository will contain a package which provides the functionalities needed to evaluate a coupon for "Coupons"

## Installation

1. Get the SDK for your project: go get -u github.com/fpinatares/couponEvaluator
2. Set the environment variable COUPONS_SERVER: export COUPONS_SERVER="http://localhost:8082"

## Usage Example
```golang
package main

import (
	"fmt"

	"github.com/fpinatares/couponEvaluator"
)

func main() {
	requiredKeys := []string{"total", "amount"}
	values := make(map[string]interface{}, 8)
	values["total"] = 100
	values["amount"] = 500
	condition := "total < amount"
	result := couponEvaluator.Evaluate(requiredKeys, values, condition)
	fmt.Println("Does the coupon applies? "+result)
}
```

> Does the coupon applies? true
