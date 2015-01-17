package main

import (
	"github.com/jcbwlkr/terminus"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	app := terminus.NewApp(termbox.ColorWhite, termbox.ColorBlack)

	menu := terminus.NewMenu(app)

	menu.Title = `
POST HYDRA
`
	menu.AddOption(&terminus.MenuOption{"About", about})
	menu.AddOption(terminus.NewExitOption("Quit"))

	menu.Run()
}

func about(app *terminus.App) int {
	app.Clear()
	app.DrawLine("Post Hydra!", 1, 1)
	app.DrawLine("This is an app!", 1, 3)
	termbox.Flush()

	time.Sleep(1 * time.Second)

	return terminus.Continue
}
