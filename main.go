package main

import (
	"fmt"

	zhun_module "github.com/ahmad-abrar-dev/go-hello-module/v2"
)

func main() {

	get_package := zhun_module.SayHello("zhun")
	fmt.Println(get_package)
}
