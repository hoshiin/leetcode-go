package leetcode

import "fmt"

type ValidWordAbbr struct {
	dict map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
	dict := map[string]string{}
	for _, s := range dictionary {
		key := ""
		count := 0
		for i, r := range s {
			if i == 0 {
				key = fmt.Sprintf("%s%s", key, string(r))
			} else if i == len(s)-1 {
				key = fmt.Sprintf("%s%d", key, count)
				key = fmt.Sprintf("%s%s", key, string(r))
			} else {
				count++
			}
		}
		val, ok := dict[key]
		if !ok {
			dict[key] = s
		} else if val != s {
			dict[key] = ""
		}
	}
	return ValidWordAbbr{dict: dict}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	key := ""
	count := 0
	for i, r := range word {
		if i == 0 {
			key = fmt.Sprintf("%s%s", key, string(r))
		} else if i == len(word)-1 {
			key = fmt.Sprintf("%s%d", key, count)
			key = fmt.Sprintf("%s%s", key, string(r))
		} else {
			count++
		}
	}
	if val, ok := this.dict[string(key)]; !ok || val == word {
		return true
	}
	return false
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */
