package main

import "fmt"


func main() {
    uri := "https://example.org:6001/v2/snacks?"
    r := "req=snickers"
    q := "quantity=1"
    s := "size=king"

    str := fmt.Sprintf("%s%s&%s&%s", uri, r, q, s)

    fmt.Printf("%s\n", str)
}
