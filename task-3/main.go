package main

import "fmt"

func main() {
	strings1 := []string{"qwe", "qweasd", "qwsdfsdf", "tr"}
	strings2 := []string{"qwe", "qweasd", "qwsdfsdf"}
	strings3 := []string{"qwe", "qeasd", "qwsdfsdf"}
	strings4 := []string{"qwe", "qweasd", "qwesdfsdf"}

	fmt.Println(mostCommonPrefix(strings1))
	fmt.Println(mostCommonPrefix(strings2))
	fmt.Println(mostCommonPrefix(strings3))
	fmt.Println(mostCommonPrefix(strings4))
}

func mostCommonPrefix(ss []string) string {
	if len(ss) == 0 {
		return ""
	}

	if len(ss) == 1 {
		return ss[0]
	}

	var prefix []byte
	var i int

loop:
	for {
		if len(ss[0]) == i {
			break
		}

		char := ss[0][i]

		for j := range ss[1:] {
			if len(ss[j]) == i {
				break loop
			}
			if char != ss[j][i] {
				break loop
			}
		}

		prefix = append(prefix, char)
		i++
	}

	return string(prefix)
}
