package graphs

/*
You are given two words, beginWord and endWord, and also a list of words wordList. All of the given words are of the same length, consisting of lowercase English letters, and are all distinct.

Your goal is to transform beginWord into endWord by following the rules:

You may transform beginWord to any word within wordList, provided that at exactly one position the words have a different character, and the rest of the positions have the same characters.
You may repeat the previous step with the new word that you obtain, and you may do this as many times as needed.
Return the minimum number of words within the transformation sequence needed to obtain the endWord, or 0 if no such sequence exists.

Example 1:

Input: beginWord = "cat", endWord = "sag", wordList = ["bat","bag","sag","dag","dot"]

Output: 4
Explanation: The transformation sequence is "cat" -> "bat" -> "bag" -> "sag".

Example 2:

Input: beginWord = "cat", endWord = "sag", wordList = ["bat","bag","sat","dag","dot"]

Output: 0
Explanation: There is no possible transformation sequence from "cat" to "sag" since the word "sag" is not in the wordList.

Constraints:

1 <= beginWord.length <= 10
1 <= wordList.length <= 100
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	wordSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	// Each side keeps only its current frontier (one distance ring) plus a
	// visited set used purely to avoid re-expanding a word on that side.
	forFrontier, backFrontier := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	forVis, backVis := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	forLevel, backLevel := 1, 1

	for len(forFrontier) > 0 && len(backFrontier) > 0 {
		if len(forFrontier) <= len(backFrontier) {
			next, met := wordLadderStep(forFrontier, forVis, backFrontier, wordSet)
			if met {
				return forLevel + backLevel
			}

			forFrontier, forLevel = next, forLevel+1
		} else {
			next, met := wordLadderStep(backFrontier, backVis, forFrontier, wordSet)
			if met {
				return forLevel + backLevel
			}

			backFrontier, backLevel = next, backLevel+1
		}
	}

	return 0
}

func wordLadderStep(front, visited, opposite, wordSet map[string]bool) (map[string]bool, bool) {
	next := map[string]bool{}
	for word := range front {
		for i := range len(word) {
			for char := byte('a'); char <= byte('z'); char++ {
				if char == word[i] {
					continue
				}

				neighbour := word[:i] + string(char) + word[i+1:]
				if opposite[neighbour] { // The two waves collided.
					return next, true
				}

				if wordSet[neighbour] && !visited[neighbour] {
					visited[neighbour] = true
					next[neighbour] = true
				}
			}
		}
	}

	return next, false
}
