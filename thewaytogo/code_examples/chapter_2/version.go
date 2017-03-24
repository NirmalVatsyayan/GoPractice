package main1

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s", runtime.Version())
}

// Output:
//go1.7.3PS
