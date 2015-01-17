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

	app := terminus.NewApp(termbox.ColorWhite, termbox.ColorBlack)

	menu := terminus.NewMenu(app)

	menu.Title = `
POST HYDRA
`
	menu.AddOption(&terminus.MenuOption{"About", about})
	// TODO this should be based on the configured readers
	menu.AddOption(&terminus.MenuOption{"Read Wild Apricot", readWildApricot})
	menu.AddOption(terminus.NewExitOption("Quit"))

	menu.Run()
}

func about(app *terminus.App) int {
	app.Clear()

	app.DrawLine("Post Hydra!", 1, 1)
	// TODO Provide better help text
	app.DrawLine("This is an app!", 1, 3)
	app.DrawLine("Press Enter to continue", 1, 5)
	termbox.Flush()

	app.WaitForEnter()

	return terminus.Continue
}

func readWildApricot(app *terminus.App) int {
	// TODO this should pull from WA and print options to the screen
	return terminus.Continue
}
