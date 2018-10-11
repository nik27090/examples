package main

import (
	"fmt"
	subpkg "github.com/polis-mail-ru-golang-1/examples/lec2/pkg/sub"
)

func main() {
	fmt.Println(subpkg.ShownConst)

	s := subpkg.Sub{Shown: 42}
	fmt.Println(s.Shown)
	fmt.Println(s.PublicGet())
}
