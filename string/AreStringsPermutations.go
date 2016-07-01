package main

import (
	"sort"
	"strings"
)

func main() {
	println(AreStringPermutation("cat", "act"))
	println(AreStringPermutation("cat", "actt"))

}

func AreStringPermutation(str1 string, str2 string) bool {
	if(len(str1) != len(str2)) {
		return false
	};
	sorted1 := SortString(str1)
	sorted2 := SortString(str2)

	return strings.EqualFold(sorted1, sorted2);
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
