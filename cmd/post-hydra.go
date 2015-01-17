package main

import (
	"github.com/jcbwlkr/terminus"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	menu := new(terminus.Menu)

	menu.Title = `
POST HYDRA
`
	menu.Options = make([]terminus.MenuOption, 1)
	menu.Options[0] = terminus.NewExitOption("Quit")

	menu.Run()

}
