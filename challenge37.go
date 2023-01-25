/* Alta3 Research | RZFeeser
   CHALLEGE 01 - Reading and displaying "n" lines of a file */


package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func main() {

    var count int
    flag.IntVar(&count, "n", 5, "number of lines to read from the file")

    var in io.Reader

    flag.Parse()

    filename := flag.Arg(0)
    if filename != "" {
        f, err := os.Open(filename)

        if err != nil {
            panic(err)
        }

        defer f.Close()

        in = f


        buf := bufio.NewScanner(in)

        for i := 0; i < count; i++ {
            if !buf.Scan() {
                break
            }

            fmt.Println(buf.Text())
        }

        if err := buf.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "error reading: err:", err)
        }


    }
}
