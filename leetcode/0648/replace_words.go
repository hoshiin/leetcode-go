package leetcode

import "strings"

func ReplaceWords(dictionary []string, sentence string) string {
	wordMap := map[string]struct{}{}
	for _, word := range dictionary {
		wordMap[word] = struct{}{}
	}
	curr := ""
	ans := ""
	skip := false
	for _, letter := range sentence {
		if letter == ' ' {
			ans = ans + curr + " "
			curr = ""
			skip = false
			continue
		}
		if skip {
			continue
		}
		curr = curr + string(letter)
		if _, ok := wordMap[curr]; ok {
			ans = ans + curr
			curr = ""
			skip = true
		}
	}
	return ans + curr
}

func ReplaceWordsTrie(dictionary []string, sentence string) string {
	t := trie{m: make(map[byte]*trie)}
	for _, word := range dictionary {
		t.insert(word)
	}
	ans := strings.Fields(sentence)
	for i, word := range ans {
		if s := t.search(word); s != "" {
			ans[i] = s
		}
	}
	return strings.Join(ans, " ")
}

type trie struct {
	m map[byte]*trie
}

func (t *trie) insert(word string) {
	current := t
	for i := 0; i < len(word); i++ {
		w := word[i]
		if v, ok := current.m[w]; ok {
			current = v
		} else {
			current.m[w] = &trie{m: make(map[byte]*trie)}
			current = current.m[w]
		}
	}
	current.m['0'] = &trie{m: make(map[byte]*trie)}
}

func (t *trie) search(word string) string {
	current, ans := t, ""
	for i := 0; i < len(word); i++ {
		v, ok := current.m[word[i]]
		if !ok {
			return ""
		}
		ans = ans + string(word[i])
		if v.m['0'] != nil {
			return ans
		}
		current = v
	}
	return ans
}
