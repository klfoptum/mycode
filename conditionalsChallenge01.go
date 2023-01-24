package main


import (
    "fmt"
    "math/rand"
    "time"
)


func init() {
    rand.Seed(time.Now().Unix())
}

func main() {
    s := []string{"A","B","C","D","E"}

    if choice := s[rand.Intn(len(s))]; choice == "A" {
        fmt.Println("A was chosen")
    } else if choice == "B" {
        fmt.Println("B was chosen")
    } else {
        fmt.Println("A and B weren't chosen")
    }
}
