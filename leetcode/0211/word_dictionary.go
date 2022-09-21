package leetcode

type WordDictionary struct {
	t *trie
}

type node struct {
	children [26]*node
	isEnd    bool
}

type trie struct {
	root *node
}

func Constructor() WordDictionary {
	t := &trie{root: &node{}}
	return WordDictionary{t}
}

func (t *trie) insert(word string) {
	curr := t.root
	for _, char := range word {
		idx := char - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &node{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (t *trie) search(word string) bool {
	ans := false
	searchHelper(word, 0, t.root, &ans)
	return ans
}

func searchHelper(word string, pos int, n *node, res *bool) {
	if n == nil {
		return
	}
	if pos == len(word) {
		if n.isEnd {
			*res = true
		}
		return
	}
	if word[pos] == '.' {
		for k := 0; k < 26; k++ {
			searchHelper(word, pos+1, n.children[k], res)
		}
	} else {
		charPos := word[pos] - 'a'
		searchHelper(word, pos+1, n.children[charPos], res)
	}
}

func (w *WordDictionary) AddWord(word string) {
	w.t.insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	return w.t.search(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
