package ex03

import (
	"os"
	"strings"
)


func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	
}


func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	
}


func Echo3() {
	strings.Join(os.Args[1:], " ")
	
}