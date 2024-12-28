package main

import (
    "fmt"
    "word_sorter_project/internal/sorter"
)

func main() {
    words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "csssssd", "fdz", "kf", "zc", "lklklklklklklkl", "l"}
    
    sorted := sorter.SortByACount(words)
    fmt.Println("Sorted words:", sorted)
}