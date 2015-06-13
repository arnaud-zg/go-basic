/*
** go-basic
**
** Made by zheng_b
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	str := "";
	fields := strings.Fields(s)
	ret := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		str = fields[i]
		(ret[str])++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
