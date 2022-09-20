package leetcode

import "strings"

type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{m: map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	ans := 0
	for key, val := range this.m {
		if strings.HasPrefix(key, prefix) {
			ans += val
		}
	}
	return ans
}

type TrieNode struct {
	children map[byte]*TrieNode
	sum      int
}

type MapSumTrie struct {
	m    map[string]int
	root *TrieNode
}

func ConstructorTrie() MapSumTrie {
	return MapSumTrie{
		m:    map[string]int{},
		root: &TrieNode{children: make(map[byte]*TrieNode)},
	}
}

func (this *MapSumTrie) Insert(key string, val int) {
	diff := val - this.m[key]
	this.m[key] = val
	cur := this.root
	for i := range key {
		_, ok := cur.children[key[i]]
		if !ok {
			trie := TrieNode{children: make(map[byte]*TrieNode)}
			cur.children[key[i]] = &trie
		}
		cur = cur.children[key[i]]
		cur.sum += diff
	}
}

func (this *MapSumTrie) Sum(prefix string) int {
	cur := this.root
	for i := range prefix {
		_, ok := cur.children[prefix[i]]
		if !ok {
			return 0
		}
		cur = cur.children[prefix[i]]
	}
	return cur.sum
}

/**
 * Your MapSumTrie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
