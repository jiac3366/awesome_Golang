package main

func isAnagram(s string, t string) bool {
	//sCount := getCount(s)
	//tCount := getCount(t)
	//return mapEqual(sCount, tCount)

	count := [26]int{}
	for _, r := range s {
		count[r-rune('a')] += 1
	}
	for _, r := range t {
		count[r-rune('a')] -= 1
	}

	return count == [26]int{}
}

func getCount(s string) map[string]int {
	count := make(map[string]int, 0)
	for _, sItem := range s {
		sItem := string(sItem)
		sNum, ok := count[sItem]
		if !ok {
			count[sItem] = 1
		} else {
			count[sItem] = sNum + 1
		}
	}
	return count
}

func mapEqual(map1 map[string]int, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value := range map1 {
		if mapV2, ok := map2[key]; !ok || value != mapV2 {
			return false
		}
	}

	return true

}

func isAnagram2(s string, t string) bool {

	record := [26]int{}
	for _, item := range s {
		record[rune(item)-rune('a')] += 1
	}

	for _, item := range t {
		record[rune(item)-rune('a')] -= 1
	}

	return record == [26]int{}
}

func main() {
	println(isAnagram2("hello", "ohell"))
}
