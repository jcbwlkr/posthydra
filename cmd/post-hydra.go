package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jcbwlkr/posthydra"
	"github.com/jcbwlkr/terminus"
	"github.com/nsf/termbox-go"
)

var config *posthydra.Config

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "PostHydra"
	cliApp.Usage = "Read posts from one source and post them to others. Like a hydra."
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "",
			Usage:  "Path to the config file",
			EnvVar: "POSTHYDRA_CONFIG",
		},
	}
	cliApp.Action = run

	cliApp.Run(os.Args)
}

func run(c *cli.Context) {
	if c.String("config") == "" {
		log.Fatalln("You must define a config with either --config or the environment variable POSTHYDRA_CONFIG")
	}
	var err error

	config, err = posthydra.NewConfig(c.String("config"))
	if err != nil {
		log.Fatalln(err)
	}

	if err := termbox.Init(); err != nil {
		log.Fatalln(err)
	}
	defer termbox.Close()

	app := terminus.NewApp(termbox.ColorWhite, termbox.ColorBlack)

	menu := terminus.NewMenu(app)

	menu.Title = `
       MMMMMMMMMMMMMMMMM?DMMOI:              INM8MMMMMMMMMMMMMMMMMM
       MMMMMMMMMMMMMMMM~                         ,NMMMMMM:MMMMMMMMM
       MMMMDNMNZMMON     ..              .8O~:ZM,   =MM.    $MMMMMM
       MMM.     7N   I:      ,D        :.        .8   +Z      MMMMM
       MM      N=   7          I.     Z            M   =MMM    NMMM
       M,   MMM.   :            .    =              Z   .MM=   MMMM
       M,   MM=   $ 8MMO         M   N MMMM,        O    $~   +$MMM
       MM   .D    MNMMMMN        N   N?MMM N        Z     M  N8MMMM
       MMM~ M     I.MMM$.        M   ..?MMM         $     ,M7MMMMMM
       MMMMZM      +            =     ?            N       MMMMMMMM
       MMMM8?       N          D =MMMM, 8         7        =ZMMMMMM
       MMMMM          O.    .$= OMMMMMM   +MZ$N8            MMMMMMM
       MMMMM                  .,  NMM~  I                   MMMMMMM
       MMMMM                 .           .,                 MMMMMMM

██████╗  ██████╗ ███████╗████████╗██╗  ██╗██╗   ██╗██████╗ ██████╗  █████╗
██╔══██╗██╔═══██╗██╔════╝╚══██╔══╝██║  ██║╚██╗ ██╔╝██╔══██╗██╔══██╗██╔══██╗
██████╔╝██║   ██║███████╗   ██║   ███████║ ╚████╔╝ ██║  ██║██████╔╝███████║
██╔═══╝ ██║   ██║╚════██║   ██║   ██╔══██║  ╚██╔╝  ██║  ██║██╔══██╗██╔══██║
██║     ╚██████╔╝███████║   ██║   ██║  ██║   ██║   ██████╔╝██║  ██║██║  ██║
╚═╝      ╚═════╝ ╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝   ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
`
	menu.AddOption(&terminus.MenuOption{"About PostHydra", about})
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
