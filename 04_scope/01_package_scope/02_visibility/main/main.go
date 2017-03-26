package main

import (
	"fmt"

	"github.com/NirmalVatsyayan/GoPractice/04_scope/01_package_scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
