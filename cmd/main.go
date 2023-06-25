package main

import (
	"flag"
	"fmt"

	"go-semantic-versioning/pkg"
)

func main () {
	branchPointer := flag.String("branch", "", "The PR branch to initate versioning from.")
	
	flag.Parse()

	fmt.Println("Starting go-semantic-versioning.")
	
	pkg.Parser("v1.0.0")
	pkg.Version(*branchPointer)
}