package main

import "fmt"

func main() {
	fmt.Println(isUniqueCharInString("asdfasdfasdfa#$%"));
	fmt.Println(isUniqueCharInString("abcdefghijklmnopqrstuvwxyz#$%"));
}

func isUniqueCharInString(str string) bool {
	var a [256]bool;
	for i := 0; i < len(str); i++ {
		if (a[str[i]]) {
			return false;
		}
		a[str[i]] = true;
	}
	return true;
}
