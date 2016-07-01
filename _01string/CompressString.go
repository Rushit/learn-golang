package main

import (
	"bytes"
	"strconv"
)
/**
Implement a method to perform basic string compression using the counts of
repeated characters. For example, the string aabcccccaaa would become
a2blc5a3. If the "compressed" string would not become smaller than the original
string, your method should return the original string
 */
func main() {
	println(CompressString("aabcccccaaa"))
	println(CompressString("aa"))
	println(CompressString("aacc"))
}

func CompressString(str string) string {

	var stringBuffer bytes.Buffer;
	var currentCharCount int = 0;

	for i := 0; i < len(str); i++ {
		if (i == 0) {
			stringBuffer.WriteByte(str[0])
			currentCharCount = 1
			continue
		}

		if (str[i - 1] == str[i]) {
			currentCharCount = currentCharCount + 1
		} else {
			stringBuffer.WriteString(strconv.Itoa(currentCharCount))
			stringBuffer.WriteByte(str[i])
			currentCharCount = 1
		}
	}
	stringBuffer.WriteString(strconv.Itoa(currentCharCount))
	if (len(stringBuffer.String()) < len(str)) {
		return stringBuffer.String()
	}
	return str;

}
