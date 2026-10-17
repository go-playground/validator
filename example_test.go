package validator_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

func ExampleValidate_Var_dirpath() {
	parent, err := os.MkdirTemp("", "validator-dirpath-")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(parent); err != nil {
			panic(err)
		}
	}()

	validate := validator.New()
	path := filepath.Join(parent, "new-directory")
	fmt.Println(validate.Var(parent, "dir") == nil)
	fmt.Println(validate.Var(path, "dir") == nil)
	fmt.Println(validate.Var(path, "dirpath") == nil)
	fmt.Println(validate.Var(path+string(os.PathSeparator), "dirpath") == nil)
	// Output:
	// true
	// false
	// false
	// true
}
