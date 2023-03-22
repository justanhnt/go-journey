package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/justanhnt/go-journey/cmd/hello"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "say hello to the world",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("say: ", hello.SayHelloToWorld(), cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "plus",
				Usage: "input a, b return sum of a and b",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() != 2 {
						log.Fatal("need input exact 2 integer")
					}
					listStrs := cCtx.Args().Slice()[:2]
					a, err := strconv.Atoi(listStrs[0])
					if err != nil {
						log.Fatalf("need input number instead of %s", listStrs[0])
					}

					b, err := strconv.Atoi(listStrs[1])
					if err != nil {
						log.Fatalf("need input number instead of %s", listStrs[1])
					}
					fmt.Printf("sum of: %d and %d is %d", a, b, a+b)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
