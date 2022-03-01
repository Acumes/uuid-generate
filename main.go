package main

import (
	"Acumes/uuid-generate/util/utils"
	"fmt"
)

func main() {
	s := make([]string, 0)
	s = append(s, "b651c56a-8f16-43a8-b2be-51bb6ad5b5e2")
	fmt.Println(utils.UuidGenerate(s))
}
