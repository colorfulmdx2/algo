package __array

import (
	"fmt"
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		hash[key] = append(hash[key], str)
		fmt.Println(hash)

	}

	result := make([][]string, 0, len(hash))
	for _, group := range hash {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for _, word := range strs {
		var count [26]int
		for _, r := range word {
			count[r-'a']++
		}
		hash[count] = append(hash[count], word)
	}

	result := make([][]string, 0)

	for _, value := range hash {
		result = append(result, value)
	}
	// make map
	// range through strs, inside range through each word
	// and count the index of char with array [26]int, rune - 'a' index++
	//add word to the map in array
	// loop through the map and settle array
	return [][]string{}
}
