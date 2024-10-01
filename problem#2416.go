package main

type TrieNode struct {
	count    int
	children [26]*TrieNode
}
type Trie struct {
	root *TrieNode
}

func (trie *Trie) insert(w string) {
	cur := trie.root
	for _, c := range w {
		i := c - 97
		if cur.children[i] == nil {
			cur.children[i] = &TrieNode{}
		}
		cur = cur.children[i]
		cur.count++
	}
}
func (trie *Trie) get_score(w string) int {
	cur := trie.root
	score := 0
	for _, c := range w {
		i := c - 97
		cur = cur.children[i]
		score += cur.count
	}
	return score
}

func sumPrefixScores(words []string) []int {
	ans := []int{}
	trie := &Trie{}
	trie.root = &TrieNode{}
	for _, w := range words {
		trie.insert(w)
	}
	for _, w := range words {
		ans = append(ans, trie.get_score(w))
	}
	return ans
}

/*
<---Hashmap solution (better space complexity, but worse time complexity)--->
func sumPrefixScores(words []string) []int {
	n := len(words)
	ans := []int{}
	score := make(map[string]int)
	for i := 0; i < n; i++ {
		s := words[i]
		for j := 1; j <= len(words[i]); j++ {
			score[s[:j]]++
		}
	}
	for i := 0; i < n; i++ {
		s := words[i]
		x := 0
		for j := 1; j <= len(words[i]); j++ {
			x += score[s[:j]]
		}
		ans = append(ans, x)
	}
	return ans
}
*/
