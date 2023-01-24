package main

import "fmt"

type Virtmach struct {
    ip string
    hostname string
    diskgb int
    ram int
}

func (v *Virtmach) UpdateRam(newRam int) {
    v.ram = newRam
}

func (v *Virtmach) UpdateDiskGB(newGB int) {
    v.diskgb = newGB
}

func main() {

    v1 := Virtmach{"127.0.0.1", "testhostname", 23, 8}
    
    fmt.Println(v1)

    v1.UpdateRam(16)
    v1.UpdateDiskGB(256)

    fmt.Println(v1)


}
