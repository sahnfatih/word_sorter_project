package sorter

// SortByACount sorts words by the number of 'a' characters in descending order
// If words have the same number of 'a' characters, they are sorted by length
func SortByACount(words []string) []string {
    result := make([]string, len(words))
    copy(result, words)

    for i := 0; i < len(result)-1; i++ {
        for j := i + 1; j < len(result); j++ {
            if shouldSwap(result[i], result[j]) {
                result[i], result[j] = result[j], result[i]
            }
        }
    }
    return result
}

// shouldSwap determines if two words should be swapped based on 'a' count and length
func shouldSwap(word1, word2 string) bool {
    count1 := countA(word1)
    count2 := countA(word2)
    
    if count1 != count2 {
        return count1 < count2
    }
    return len(word1) < len(word2)
}

// countA counts the number of 'a' characters in a word
func countA(word string) int {
    count := 0
    for _, char := range word {
        if char == 'a' {
            count++
        }
    }
    return count
}