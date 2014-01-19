package main

import (
	"fmt"
	"bytes"
)

func main() {
	nl := "\n"
	bnl := []byte("\n")
	bsnl := []byte(nl)
	r := bytes.Compare(bnl, bsnl)
	fmt.Println("bytes.Compare",r)
	
	a := byte('\n')
	b := byte('\n')
	fmt.Println(a == b)
	
}