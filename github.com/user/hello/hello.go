package main

import (
	"fmt"
	"go-practice/github.com/user/stringutil"
)

func main() {
	str := "Hello World!"
	//s := [] rune(str)
	//fmt.Printf("%d %c\n", len(s) , s[0])
	//for kk,ss := range s {
	//	fmt.Printf("(%d %c)\n", kk,ss)
	//}

	str1 := []rune(str)
	for kkk, sss := range str1 {
		fmt.Printf("(%d %c)\n", kkk, sss)
	}
	fmt.Printf(stringutil.Reverse("!OG, olleH"))
}
