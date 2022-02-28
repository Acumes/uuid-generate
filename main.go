package main

import (
	"Acumes/uuid-generate/util/utils"
	"fmt"
)

func main() {
	s := make([]string,0)
	s = append(s,"1")
	fmt.Println(utils.UuidGenerate(s))
}
