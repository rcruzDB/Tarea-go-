package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", commaSigned(os.Args[i]))
	}
}


func commaSigned(s string) string {
	var start, end int

	if strings.HasPrefix(s, "-") {
		start = 1
	} else {
		start = 0
	}

	if strings.Contains(s, ".") {
		end = strings.Index(s, ".")
	} else {
		end = len(s)
	}

	return s[:start] + comma(s[start:end]) + s[end:]
}


func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}