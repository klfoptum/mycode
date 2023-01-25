package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    argLength := len(os.Args[1:])    
    fmt.Printf("Arg length is %d", argLength) 
}

