package main

import (
    "fmt"
    "sort"
)

// In order to sort by a custom function in Go we need a corresponding type
// byLength type is just an alias for the builtin []string type
type byLength []string


// Implement our custom sort by converting the original fruits slice to byLength, 
// and then use "sort.Sort" on that typed slice
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
