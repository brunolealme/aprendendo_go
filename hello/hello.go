package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	var result string
	result = morestrings.ReverseRunes("Bruno")

	fmt.Println(morestrings.ReverseRunes("Bruno"))
	fmt.Println(cmp.Diff(result, result))
}
