package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sSlice := strings.Fields(s)
	r := make(map[string]int)
	for i:= 0; i < len(sSlice); i++{
		r[sSlice[i]]++
	}

	return r
}

func main() {
	wc.Test(WordCount)
}

//https://go-tour-ru-ru.appspot.com/moretypes/23