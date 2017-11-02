package main

import "fovl/CLI"
import "fmt"

func main() {
	parser := CLI.CreateArgumentParser()
	res, _ := parser.Parse()
	fmt.Println(res)
}
