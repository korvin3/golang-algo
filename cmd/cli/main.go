package main

import (
	"fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "help",
				Aliases: []string{"h", ""},
				Usage: "help",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("help")
					return nil
				},
			},
			{
				Name: "merge-sort",
				Aliases: []string{"ms"},
				Usage: "merge sort",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("help")
					return nil
				},
			},
		},
        Name:  "golang-cli-app",
        Usage: "some functions and helpers as a cli application",
        Action: func(*cli.Context) error {
            fmt.Println("boom! I say!")
            return nil
        },
    }
	
	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
