package validator_test

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ExampleValidate_Var_numberAndNumeric() {
	validate := validator.New()

	for _, value := range []any{"1", "-1", "1.618", 1.618, "1e3"} {
		fmt.Printf("%#v: number=%t, numeric=%t\n", value,
			validate.Var(value, "number") == nil,
			validate.Var(value, "numeric") == nil)
	}

	// Output:
	// "1": number=true, numeric=true
	// "-1": number=false, numeric=true
	// "1.618": number=false, numeric=true
	// 1.618: number=true, numeric=true
	// "1e3": number=false, numeric=false
}
