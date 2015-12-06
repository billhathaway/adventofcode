package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strings"
)

func findHash(key string) int {
	var num int
	var buf bytes.Buffer
	prefix := "00000"
	for {
		buf.WriteString(key)
		buf.WriteString(fmt.Sprintf("%d", num))
		data := md5.Sum(buf.Bytes())
		hex := fmt.Sprintf("%x", data)
		if strings.HasPrefix(hex, prefix) {
			return num
		}
		buf.Reset()
		num++
	}
}

func main() {
	fmt.Printf("%d\n", findHash("yzbqklnj"))
}
