package sorter

// SortByACount sorts words by the number of 'a' characters in descending order
// If words have the same number of 'a' characters, they are sorted by length
func SortByACount(words []string) []string {
    // Create a copy of the input slice to avoid modifying the original
    result := make([]string, len(words))
    copy(result, words)

    // Sort using custom comparison function
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
    
   // If the 'a' numbers are different, the larger one comes first
    if count1 != count2 {
        return count1 < count2 // We change the smaller one to be in descending order
    }
    
    // If the numbers 'a' are the same and the lengths are different
    if len(word1) != len(word2) {
        return len(word1) < len(word2) // Longest one comes first
    }
    
   // If 'a' numbers and lengths are the same, sort alphabetically
    return word1 > word2
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