package sorter

import (
    "reflect"
    "testing"
)

func TestSortByACount(t *testing.T) {
    tests := []struct {
        name     string
        input    []string
        expected []string
    }{
        {
            name:     "Sample test case",
            input:    []string{"aaaasd", "a", "aab", "aaabcd", "ef", "csssssd", "fdz", "kf", "zc", "lklklklklklklkl", "l"},
            expected: []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklkl", "csssssd", "fdz", "ef", "kf", "zc", "l"},
        },
        {
            name:     "Empty slice",
            input:    []string{},
            expected: []string{},
        },
        {
            name:     "Single word",
            input:    []string{"hello"},
            expected: []string{"hello"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := SortByACount(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("SortByACount() = %v, want %v", result, tt.expected)
            }
        })
    }
}