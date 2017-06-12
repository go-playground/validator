package main

import "github.com/gin-gonic/gin/binding"

func main() {

	binding.Validator = new(defaultValidator)

	// regular gin logic
}
