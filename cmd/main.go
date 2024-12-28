package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "word_sorter_project/internal/sorter"
)

func main() {
    var words []string

    // Check if arguments are provided
    if len(os.Args) > 1 {
        // Use command line arguments
        words = os.Args[1:]
    } else {
        // No arguments, read from standard input
        fmt.Println("Enter words (one per line, press Ctrl+D [Unix] or Ctrl+Z [Windows] to finish):")
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            word := strings.TrimSpace(scanner.Text())
            if word != "" {
                words = append(words, word)
            }
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
            os.Exit(1)
        }
    }

    // Check if we have any words to process
    if len(words) == 0 {
        fmt.Println("No words provided. Please provide words either as arguments or through standard input.")
        fmt.Println("Usage examples:")
        fmt.Println("1. As arguments: go run cmd/main.go word1 word2 word3")
        fmt.Println("2. Through input: go run cmd/main.go (then type words)")
        os.Exit(1)
    }

    // Sort the words
    sorted := sorter.SortByACount(words)

    // Print results
    fmt.Println("\nResults:")
    fmt.Println("Input:", strings.Join(words, " "))
    fmt.Println("Sorted:", strings.Join(sorted, " "))
}