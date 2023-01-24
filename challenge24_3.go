package main

import "fmt"

type Player struct {
    Lives int
    Stage int
    Inventory []string
}

func (p *Player) Greenmushroom() {
    p.Lives++
}

func (p *Player) Pickup(item string) {
    p.Inventory = append(p.Inventory, item)
}

func (p Player) CanWhistle() bool {
    return p.Stage >= 5
}

func main() {

    mario := Player{Lives: 3, Stage: 1, Inventory: []string{"Mushroom"}}

    fmt.Println(mario)

    mario.Greenmushroom()
    mario.Pickup("Cape")

    fmt.Println(mario.CanWhistle())
    mario.Stage=6
    fmt.Println(mario.CanWhistle())

    fmt.Println(mario)


}
