package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(IsRotationString("waterbottle", "bottlewater"))
	fmt.Println(IsRotationString("waterbottle", "bottlewater1"))

}

func IsRotationString(s1 string, s2 string) bool {
	s1s1 := s1 + s1
	return strings.Contains(s1s1, s2)
}
