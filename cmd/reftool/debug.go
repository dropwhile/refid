package main

import (
	"fmt"
	"strconv"
)

func PrintBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		fmt.Printf("%08s ", strconv.FormatInt(int64(b[i]), 2))
		if i != 0 && (i+1)%4 == 0 {
			fmt.Println()
		}
	}
}
