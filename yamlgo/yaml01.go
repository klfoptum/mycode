package main

import (
    "fmt"
    "io"
    "log"
    "os"

    "gopkg.in/yaml.v3"
)

func main() {

    fs, err := os.Open("startrek.yaml")
    yfile, err := io.ReadAll(fs)

    if err != nil {
        log.Fatal(err)
    }

    data := make(map[any]any)

    err2 := yaml.Unmarshal(yfile, &data)

    if err2 != nil {
        log.Fatal(err2)
    }

    // k - Key from the map (left side)
     // v - Value from the map (right side)
     for k, v := range data {
          fmt.Printf("%s -> %d\n", k, v)
     }
}
