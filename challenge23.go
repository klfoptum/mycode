
package main

import "fmt"

type Astro struct {
    name string
    age int
    mission string
    isneeded bool
}

type nasaMission struct {
    people []Astro
    number int
    message string
}

func main() {

    ast1 := Astro{"A", 3, "Eating", false}
    ast2 := Astro{"B", 5, "Sleeping", true}
    ast3 := Astro{"C", 7, "SpaceX Crew-3", false}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)

    p := []Astro{ast1, ast2, ast3}

    fmt.Println(p)

    fmt.Println(p[2].mission)

    s := nasaMission{p, len(p), "success"}
    fmt.Println(s)
    fmt.Printf("%+v", s)

}
