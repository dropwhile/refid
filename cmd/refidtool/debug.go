package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dropwhile/refid"
)

func PrintBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		fmt.Printf("%08s ", strconv.FormatInt(int64(b[i]), 2))
		if i != 0 && (i+1)%4 == 0 {
			fmt.Println()
		}
	}
}

func PrintRefID(xID refid.RefID) {
	tx := xID.Time().UTC()
	fmt.Printf("native enc:   %s\n", xID.String())
	fmt.Printf("hex enc:      %s\n", xID.ToHexString())
	fmt.Printf("base64 enc:   %s\n", xID.ToBase64String())
	fmt.Printf("tag value:    %d\n", xID.Tag())
	fmt.Printf("type:         %s\n", xID.Type())
	fmt.Printf("time(string): %s\n", tx.Format(time.RFC3339Nano))
	fmt.Printf("time(millis): %d\n", tx.UnixMilli())
}
